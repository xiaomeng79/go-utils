//文件和路径的判断
package path

import (
	"os"
)

//判断文件是否存在
// 存在 true,nil
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


//打开文件，追加模式，如果不存在就创建文件
func OpenFileOrCreate(filename string) (*os.File, error) {

	return os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm|os.ModeTemporary)
}




