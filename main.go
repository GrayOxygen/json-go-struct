package main

import (
	"fmt"
	"github.com/GrayOxygen/json-go-struct/apis"
	"github.com/GrayOxygen/json-go-struct/util"
	_ "github.com/robertkrimen/otto/underscore"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {

	//读取json文件	START
	fmt.Println(util.GetCurPath())
	bytes, err := ioutil.ReadFile("./testjson/test3.json")

	jsonStr := string(bytes)
	////读取json文件  END
	util.Log.Printf("\n 传入的json为 %s   \n", jsonStr)
	nestStructStr, structStr, err := apis.JSON2Struct(jsonStr);
	fmt.Println("Start............===============")
	fmt.Println(nestStructStr)
	fmt.Println("===============")
	fmt.Println(structStr)
	fmt.Println("===============")
	if err != nil {
		fmt.Println(err)
	}
	util.Log.Printf("测试 v%s pid=%d started with processes: %d", "1.0.0", os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))
}
