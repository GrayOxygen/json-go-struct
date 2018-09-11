package apis

import (
	"github.com/golang-collections/collections/stack"
	"os"
	"github.com/GrayOxygen/json-go-struct/errors"
	"github.com/GrayOxygen/json-go-struct/parser"
	"github.com/GrayOxygen/json-go-struct/tree"
	"github.com/GrayOxygen/json-go-struct/util"
	"strings"

	"fmt"
)

var (
	nestStrutPath = util.GetCurPath() + "/嵌套Struct"
	structName    = "StructName" //输出的struct名称(顶级)  TODO 支持自定义
	//左大括，成对出栈找到子struct
	leftStack = stack.New()
)
//nestStruct struct
func JSON2Struct(jsonStr string) (string, string, error) {
	util.Log.Printf("\n 传入的json为 %s   \n", jsonStr)
	//1，解析json为嵌套struct字符串
	nestStructStr, err := parser.JsonToNestStruct(jsonStr)

	if err != nil {
		fmt.Println("解析json为嵌套struct字符串失败：：：", err)
		return "", "", err
	}
	if nestStructStr == "" {
		fmt.Println("json解析失败，请检查格式是否正确，常见错误有：json中带有//注释，中英文的逗号等")
		return "", "", errors.New("json解析失败，请检查格式是否正确，常见错误有：json中带有//注释，中英文的逗号等")
	}
	if !strings.Contains(nestStructStr, "{") || ! strings.Contains(nestStructStr, "}") {
		return "", "", errors.New("json解析失败，请检查格式是否正确，常见错误有：json中带有//注释，中英文的逗号等")
	}

	if util.Exists(nestStrutPath) {
		if err := os.Remove(nestStrutPath); err != nil {
			fmt.Println("删除 嵌套文件失败：：：", err)
			return "", "", err
		}
	}
	if err := util.WriteTrunc(nestStrutPath, nestStructStr); err != nil {
		fmt.Println("内嵌Struct结构写入文件失败：：：", err)
		return "", "", err
	}

	root := &tree.TreeNode{}
	root.Level = 0 //层级为0，返回的数据才是真的tree，层级从1开始
	lineCount := 1
	nss := strings.Split(nestStructStr, "\n") //下标0是第一行，1是第二行...
	util.Log.Printf("\n 直接按行划分，数组为 %s   \n", nss)

	parser.ClearCache()
	root, _ = parser.DFS(nss, leftStack, lineCount, root)
	//打印树
	util.PrintTree(root)

	//3，遍历树，输出非嵌套struct到文件中
	res, err := parser.Generate(root)
	util.Log.Printf("\n 返回的非嵌套struct为 \n %s   \n", res)
	if err != nil {
		fmt.Println("生成非嵌套文件失败：：：", err)
		return "", "", err
	}
	return nestStructStr, res, nil

}
