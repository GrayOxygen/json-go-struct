package util

import (
	"go/format"
	"os"
	"os/exec"
	"path/filepath"
	"bufio"
)

//追加写入文件
func WriteAppend(filePath string, str string) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	//f.Write([]byte(str ))
	//f.WriteString("\n")
	//格式化后写入
	content, err := format.Source([]byte(str + "\n"))
	if err != nil {
		return err
	}
	f.Write(content)
	return nil

}

//覆盖写入
func WriteTrunc(path string, str string) error {
	//写入文件，内嵌的struct结构字符串
	f, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	//不存在则创建
	if err != nil && !os.IsExist(err) {
		f, err = os.Create(path)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}
	defer f.Close()
	f.Write([]byte(str ))
	f.WriteString("\n")
	return nil
}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/*获取当前文件执行的路径*/
func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	return rst
}

//按行读取文件
func FileToLines(filePath string) (lines []string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}
