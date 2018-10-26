package parser

import (
	"fmt"
	"github.com/GrayOxygen/json-go-struct/enums"
	"github.com/GrayOxygen/json-go-struct/tree"
	"github.com/GrayOxygen/json-go-struct/util"
	"go/format"
	"os"
	"strings"
)

func GeneratorProto(node *tree.TreeNode) (string, error) {
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

	//===========================================================================
	////先得到非内潜的struct，再调整为proto文件
	//bytes, err := ioutil.ReadFile("./testjson/test3.json")
	//
	//jsonStr := string(bytes)
	//////读取json文件  END
	//nestStructStr, structStr, err := apis.JSON2Struct(jsonStr)
	//
	//fmt.Println("Start............===============")
	//fmt.Println(nestStructStr)
	//fmt.Println("===============")
	//fmt.Println(structStr)
	//fmt.Println("===============")
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func generateChildrenProto(node *tree.TreeNode) error {
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

//TODO
func genOneNodeProto(node *tree.TreeNode) error {
	temp := node.Value.DefineInfo
	objs := tree.GetSonObjs(node)
	for _, c := range objs { //替换id为规范属性声明
		if c.Type == enums.STRUCT {
			temp = strings.Replace(temp, c.Id, c.Name+" "+c.Name, -1)
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
