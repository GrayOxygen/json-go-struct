package util

import (
	"strings"
	"bufio"
)

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

//按行读取字符串
func StringToLines(s string) (lines []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

//移除字符串中的空行
func RemoveEmptyLineString(s string) string {
	var lines string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		temp := scanner.Text()
		if strings.TrimSpace(temp) != "" {
			lines += temp + "\n"
		}
	}
	if len(lines) < 1 {
		return s
	}

	return lines
}
