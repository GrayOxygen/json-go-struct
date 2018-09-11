package model


import "github.com/GrayOxygen/json-go-struct/enums"



type StructObj struct {
	Id         string
	Name       string
	Type       enums.PropertyType
	DefineInfo string
	Describe   string
}

//解析过的括号
type UsedBrace struct {
	Start int
	End   int
}

//左括号
type Brace struct {
	B string
	N int
}
