package fileutils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"log"
	"strings"
)

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func SaveData(fileName string, data string, isAppend bool) error {
	i := os.O_CREATE | os.O_WRONLY
	if isAppend {
		i = i | os.O_APPEND
	}
	f, err := os.OpenFile(fileName, i, 0666)
	if err != nil {
		return err
	}

	defer f.Close()
	f.WriteString(data)
	return nil
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

//func SaveFile (fileName string, data string) error{
//
//}

//获取文件大小
func GetFileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}
