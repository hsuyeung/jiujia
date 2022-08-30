package fileutil

import (
	"io/ioutil"
	"os"
)

// 文件相关工具类

// IsExist 判断一个路径或文件是否存在
func IsExist(dirOrFileName string) (isExist bool) {
	_, err := os.Stat(dirOrFileName)
	return err == nil || os.IsExist(err)
}

// ReadAllText 从文件中读取所有文本数据
func ReadAllText(path string) (text string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(bytes)
}
