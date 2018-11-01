package parser

import (
	"fmt"
	"github.com/GrayOxygen/json-go-struct/enums"
	"github.com/GrayOxygen/json-go-struct/tree"
	"github.com/GrayOxygen/json-go-struct/util"
	"os"
	"strconv"
	"strings"
)

var (
	obj_unique = "obj-unique-by-davidwang"
)

func GenerateProto(node *tree.TreeNode) (string, error) {
	//文件存在则删除
	if util.Exists(structPath) {
		if err := os.Remove(structPath); err != nil {
			fmt.Println("删除非嵌套文件失败：：：", err)
			return "", err
		}
	}

	//同名的名字作区分
	distinctStructName(node)

	err := genOneProtoNode(node)
	//util.Log.Printf("\n genOneProtoNode root 参数是 %s \n 后结果是 %s   \n", node, resStruct)

	err = generateProtoChildren(node)
	if err != nil {
		return resStruct, err
	}
	return resStruct, nil
}

func generateProtoChildren(node *tree.TreeNode) error {
	for _, item := range node.Children {
		err := genOneProtoNode(item)
		if err != nil {
			return err
		}
		//util.Log.Printf("\n generateProtoChildren 结果是 %s   \n 参数为 %s\n", resStruct, item)
		err = generateProtoChildren(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func genOneProtoNode(node *tree.TreeNode) error {
	temp := node.Value.DefineInfo
	temp = strings.Replace(temp, "struct", " ", -1)
	objs := tree.GetSonObjs(node)

	// 对象和数组的处理
	for _, c := range objs { //替换id为规范属性声明
		if c.Type == enums.STRUCT {
			temp = strings.Replace(temp, c.Id, "    "+c.Name+" "+util.LowerFirstRune(c.Name)+" "+obj_unique, -1)
		}
		if c.Type == enums.ARRAY {
			temp = strings.Replace(temp, c.Id, "    repeated "+c.Name+" "+util.LowerFirstRune(c.Name)+" "+obj_unique, -1)
		}
	}
	//不是自己直接儿子进行忽略
	objs = tree.GetSonChildrenObjs(node)
	content := make([]byte, 0)
	ct := temp
	var err error
	for _, c := range objs { //子的孩子，无需设置属性，忽略掉
		temp = strings.Replace(temp, c.Id, "", -1)
	}

	//必须格式化
	content = []byte("message " + temp + "\n")
	if err != nil {
		return err
	}
	ct = util.RemoveEmptyLineString(string(content))

	//调整位置
	ct = adjustDefineInfo(ct)

	//增加proto规范的内容
	final_res := ""

	scanner := strings.Split(ct, "\n") //下标0是第一行，1是第二行...
	lineCount := 2
	final_res += scanner[0] + "\n"
	for index := lineCount - 1; index < len(scanner)-1; index++ { //不管第一行和最后一行
		if strings.Contains(scanner[index], obj_unique) {
			scanner[index] = strings.Replace(scanner[index], obj_unique, " \n", -1)
			scanner[index] = strings.TrimRight(scanner[index], " ")
			scanner[index] = strings.TrimRight(scanner[index], "\n")
			scanner[index] = strings.TrimRight(scanner[index], "\t")
		}
		if strings.Contains(strings.TrimSpace(scanner[index]), "}") {
			final_res += "}"
			break
		}
		if strings.TrimSpace(scanner[index]) != "" { //去掉json描述，更改类型和变量名的位置
			final_res += scanner[index] + "=" + strconv.Itoa(index) + ";" + "\n"
		}
	}
	final_res += scanner[len(scanner)-1] + "\n"

	//再去掉空行
	//ct = strings.Replace(string(content), "\n\n", "", -1)

	//util.Log.Printf("\n genOneNode::后的内容：：：    %s\n", content)
	resStruct += string(final_res)
	//util.Log.Printf("\n genOneNode::后的resStruct内容：：：    %s\n", resStruct)
	return nil
}

//调整define info信息，保证字段的定义符合proto规范
func adjustDefineInfo(ct string) string {
	final_res := ""

	scanner := strings.Split(ct, "\n") //下标0是第一行，1是第二行...
	lineCount := 2
	final_res += scanner[0] + "\n"
	for index := lineCount - 1; index < len(scanner); index++ { //不管第一行和最后一行
		if strings.Contains(strings.TrimSpace(scanner[index]), "}") {
			final_res += "}\n"
			break
		}
		if strings.TrimSpace(scanner[index]) != "" { //去掉json描述，更改类型和变量名的位置
			if strings.Contains(scanner[index], obj_unique) { //存在对象类型的不用再颠倒了
				final_res += scanner[index] + "\n"
				continue
			}
			arrays := strings.Split(scanner[index], " ")
			//颠倒为 类型  变量名	这样的顺序
			if strings.Contains(arrays[1], "float") {
				arrays[1] = "float"
			}
			if strings.Contains(arrays[1], "int") {
				arrays[1] = "int32"
			}
			if strings.Contains(arrays[1], "bool") {
				arrays[1] = "bool"
			}
			final_res += "	" + strings.TrimSpace(arrays[1]) + " " + util.LowerFirstRune(strings.TrimSpace(arrays[0])) + "\n"
		}
	}
	return final_res

}
