package util

import "strings"

//  首字母小写
func LowerFirstRune(str string) string {
	var res string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			res += strings.ToLower(string(vv[i])) // + string(vv[i+1])
		} else {
			res += string(vv[i])
		}
	}

	return res
}

//  首字母大写
func UpperFirstRune(str string) string {
	var res string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			res += strings.ToUpper(string(vv[i])) // + string(vv[i+1])
		} else {
			res += string(vv[i])
		}
	}

	return res
}
