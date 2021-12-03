package file

import "os"

//判断给定的路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
// 如果不存在，也是 false
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
// 同样使用 IsDir 如果不存在的也是 false, 因为这个地方只设计了一个返回值
func IsFile(path string) bool {
	return !IsDir(path)
}
