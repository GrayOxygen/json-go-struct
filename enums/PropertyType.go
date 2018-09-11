package enums

//属性的类型
type PropertyType int

const (
	BASIC  PropertyType = iota // 基础类型 0
	STRUCT                     //1
	ARRAY                      // 2
)
