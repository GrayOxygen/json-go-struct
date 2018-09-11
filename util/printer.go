package util

import (
	"fmt"
	"github.com/GrayOxygen/json-go-struct/tree"
	"reflect"
)

func PrettyPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("获取到数据:")
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}
func PrintTree(node *tree.TreeNode) {
	Log.Printf("\n 当前层级 %d \n 打印非嵌套树结构 \n %s   \n", node.Level, node.Value.DefineInfo)
	printTreeChildren(node)
}

func printTreeChildren(node *tree.TreeNode) {
	for _, item := range node.Children {
		Log.Printf("\n 当前层级 %d \n %s \n", item.Level, item.Value.DefineInfo)
		printTreeChildren(item)
	}
}
