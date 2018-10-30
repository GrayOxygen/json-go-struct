package model

import "github.com/GrayOxygen/json-go-struct/enums"

type StructObj struct {
	Id         string
	Name       string
	Type       enums.PropertyType
	DefineInfo string //struct定义信息，如	ShippingAddressInfo struct {
	Describe   string //如	BuyerID   string `json:"buyerId"`的`json:"buyerId"`
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
