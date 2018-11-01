package parser

import (
	"github.com/GrayOxygen/json-go-struct/consts"
	"github.com/GrayOxygen/json-go-struct/enums"
	"github.com/GrayOxygen/json-go-struct/model"
	"github.com/GrayOxygen/json-go-struct/tree"
	"github.com/GrayOxygen/json-go-struct/util"
	"github.com/golang-collections/collections/stack"
	"github.com/robertkrimen/otto"
	"strings"

	"bytes"
	"fmt"
	"go/format"
	"math"
	"os"

	"github.com/google/uuid"
)

var (
	structPath    = util.GetCurPath() + "/非嵌套Struct.go"
	jsPath        = util.GetCurPath() + "/json-to-go.js"
	nestStructStr = ""
	structName    = "Object" //输出的struct名称(顶级)

	maxLevel    = 0                                //树形结构最大层级
	defineInfos = make(map[int]model.StructObj, 0) //key：起始行号，value：对应的struct，用id标记struct属性
	usedBrace   = make([]*model.UsedBrace, 0)      //读取过的struct，替换为id
	resStruct   = ""
)

//清除缓存
func ClearCache() {
	maxLevel = 0                                   //树形结构最大层级
	defineInfos = make(map[int]model.StructObj, 0) //key：起始行号，value：对应的struct，用id标记struct属性
	usedBrace = make([]*model.UsedBrace, 0)        //读取过的struct，替换为id
	resStruct = ""
}

//dfs深度优先遍历，解析struct为树
func DFS(fileScanner []string, leftStack *stack.Stack, lineCount int, root *tree.TreeNode) (resNode *tree.TreeNode, resCount int) {
	//从第lineCount一直往后读
	for index := lineCount - 1; index < len(fileScanner); index++ {
		text := fileScanner[index]
		//struct行入栈
		if strings.Contains(text, "struct") && strings.Contains(text, "{") {
			b := model.Brace{B: "{", N: index + 1}
			leftStack.Push(b)
			//层级递增
			temp := tree.TreeNode{
				Level: root.Level + 1,
			}
			//递归读取
			n, c := DFS(fileScanner, leftStack, index+1+1, &temp)

			root.Children = append(root.Children, n)
			maxLevel = int(math.Max(float64(maxLevel), float64(n.Level)))
			index = c - 1
			continue
		}
		//从}层级最深的括号对依次向外开始解析
		if strings.Contains(text, "}") && !strings.Contains(text, "interface") {
			leftBrace, _ := (leftStack.Pop()).(model.Brace)

			//读取单个struct，排除自身已经读取过的struct属性，用唯一标识占位
			oneStruct := ReadStructObj(leftBrace.N, index+1, fileScanner, defineInfos)
			//读取过的struct行，记录起始结尾行号
			usedBrace = append(usedBrace, &model.UsedBrace{Start: leftBrace.N, End: index + 1})
			resNode = new(tree.TreeNode)
			resNode.Value = &oneStruct
			resNode.Level = root.Level
			resNode.Children = append(resNode.Children, root.Children...)

			return resNode, index + 1
		}
	}

	//总层级数
	root.Children[0].MaxLevel = maxLevel

	return root.Children[0], lineCount
}

//读取单个struct信息
func ReadStructObj(start, end int, array []string, defineInfos map[int]model.StructObj) model.StructObj {
	//桶映射：需读取的行号范围
	buckets := make([]int, end-start+1)

	//排除不需读取的struct属性，占位标记
	for _, item := range usedBrace { //
		if item.Start >= start && item.End <= end {
			for i := item.Start; i <= item.End; i++ {
				buckets[i-start] = 1
			}
		}
	}
	util.Log.Printf("\n 从第%d行 读取到 第%d行   \n", start, end)
	util.Log.Printf("\n 当前已记录的struct信息为   \n", defineInfos)
	util.Log.Printf("\n 占位后的桶   \n", buckets)

	dest := make([]string, 0)
	destDescribe := ""
	theType := enums.STRUCT
	first := false
	name := ""
	for i, item := range buckets {
		//需读取的行
		if item == 0 {
			if !first && strings.Contains(array[i+start-1], "[]") { //读取struct数组声明
				name = strings.Split(array[i+start-1], "[]")[0]
				array[i+start-1] = strings.Replace(array[i+start-1], "[]", "", -1)
				first = true
				theType = enums.ARRAY
			}

			if !first { //读取struct声明
				name = strings.Split(array[i+start-1], "struct")[0]
				//去掉type
				if strings.Contains(name, "type") {
					name = strings.Split(name, "type")[1]
				}
				first = true
			}

			if strings.Contains(array[i+start-1], "}") {
				destDescribe = strings.TrimSpace(strings.Split(array[i+start-1], "}")[1])           //json等描述
				dest = append(dest, strings.TrimSpace(strings.Split(array[i+start-1], "}")[0])+"}") //去掉json等描述
			} else {
				if strings.Contains(array[i+start-1], "type") && strings.Contains(array[i+start-1], "{") {
					//去掉最外层定义的type，仅在顶级struct的第一行定义需要替换type
					dest = append(dest, strings.Replace(array[i+start-1], "type", "", -1))
				} else {
					dest = append(dest, array[i+start-1])
				}
			}
		}

		//struct类型替换，只替换直接第一层的  usedBrace
		if item == 1 {
			if item, ok := defineInfos[i+start-1]; ok {
				dest = append(dest, item.Id)
			}
		}
	}

	//行号记录内嵌struct，用于后续替换
	id := uuid.New()
	temp := model.StructObj{
		Id:         id.String(),
		Name:       strings.TrimSpace(name),
		Type:       theType,
		DefineInfo: strings.Join(dest, "\n"),
		Describe:   destDescribe,
	}
	defineInfos[start] = temp

	return temp
}

//替换struct属性，同名的，从最外层往里，用父级名字附加在尾部，最里层不添加
func Generate(node *tree.TreeNode) (string, error) {
	//文件存在则删除
	if util.Exists(structPath) {
		if err := os.Remove(structPath); err != nil {
			fmt.Println("删除非嵌套文件失败：：：", err)
			return "", err
		}
	}

	//同名的名字作区分
	distinctStructName(node)
	//if node.Level == 1 {
	err := genOneNode(node)
	util.Log.Printf("\n genOneNode root 参数是 %s \n 后结果是 %s   \n", node, resStruct)
	//}
	err = generateChildren(node)
	if err != nil {
		return resStruct, err
	}
	return resStruct, nil
}

//   struct重名问题
func distinctStructName(node *tree.TreeNode) {
	nodes := tree.GetChildren(node)
	//hash  key:struct name value :node
	distinctNodes := make(map[string]*tree.TreeNode)

	for _, item := range nodes {
		if n, ok := distinctNodes[item.Value.Name]; !ok {
			distinctNodes[item.Value.Name] = n
		} else {
			fatherNode := tree.GetFatherNode(node, item)
			oldName := item.Value.Name
			newName := fatherNode.Value.Name + oldName
			item.Value.Name = newName
			item.Value.Describe = strings.Replace(item.Value.Describe, util.GetTagName(oldName), util.GetTagName(newName), -1)
			item.Value.DefineInfo = strings.Replace(item.Value.DefineInfo, oldName, newName, -1)
			distinctNodes[item.Value.Name] = n
		}
	}
}

func generateChildren(node *tree.TreeNode) error {
	for _, item := range node.Children {
		err := genOneNode(item)
		if err != nil {
			return err
		}
		util.Log.Printf("\n generateChildren递归一次 结果是 %s   \n 参数为 %s\n", resStruct, item)
		err = generateChildren(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func genOneNode(node *tree.TreeNode) error {
	temp := node.Value.DefineInfo
	objs := tree.GetSonObjs(node)
	for _, c := range objs { //替换id为规范属性声明
		if c.Type == enums.STRUCT {
			temp = strings.Replace(temp, c.Id, c.Name+" *"+c.Name+c.Describe, -1)
		}
		if c.Type == enums.ARRAY {
			temp = strings.Replace(temp, c.Id, c.Name+" []*"+c.Name+c.Describe, -1)
		}
	}
	objs = tree.GetSonChildrenObjs(node)
	content := make([]byte, 0)
	ct := temp
	var err error
	for _, c := range objs { //子的孩子，无需设置属性
		temp = strings.Replace(temp, c.Id, "", -1)
	}

	//必须格式化
	content, err = format.Source([]byte("type	 " + temp + "\n"))
	if err != nil {
		return err
	}
	ct = util.RemoveEmptyLineString(string(content))

	//再去掉空行
	//ct = strings.Replace(string(content), "\n\n", "", -1)

	util.Log.Printf("\n genOneNode::后的内容：：：    %s\n", content)
	resStruct += string(ct)
	util.Log.Printf("\n genOneNode::后的resStruct内容：：：    %s\n", resStruct)
	return nil
}

//JS解析json为嵌套struct
func JsonToNestStruct(jsonStr string) (string, error) {
	vm, err := initParser()
	if err != nil {
		return "", err
	}

	//转换参数，调用API解析为struct字符串
	jsa, err := vm.ToValue(jsonStr)
	if err != nil {
		return "", err
	}
	jsb, err := vm.ToValue(structName)
	if err != nil {
		return "", err
	}
	result, err := vm.Call("jsonToGo", nil, jsa, jsb)

	if err != nil {
		return "", err
	}

	x, _ := result.Object().Get("go")
	structStr, _ := x.ToString()

	return structStr, nil
}

//初始化json to go 的js解析器:读取源码
func initParser() (*otto.Otto, error) {
	//初始化js
	vm := otto.New()
	if _, err := vm.Run(consts.JsSourceStr); err != nil {
		return nil, err
	}
	return vm, nil
}

//  废弃:读取js-json-go.js文件来初始化json to go 的js解析器
func initParserBak() (*otto.Otto, error) {
	f, err := os.Open(jsPath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	buff := bytes.NewBuffer(nil)

	if _, err := buff.ReadFrom(f); err != nil {
		return nil, err
	}

	//初始化js
	vm := otto.New()
	temp := buff.String()
	fmt.Println(temp)
	if _, err := vm.Run(buff.String()); err != nil {
		return nil, err
	}
	return vm, nil
}
