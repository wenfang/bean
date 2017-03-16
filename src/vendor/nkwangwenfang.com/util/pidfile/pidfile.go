package pidfile

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Save 保存当前的PID到文件中
func Save(filename string) error {
	return ioutil.WriteFile(filename, []byte(fmt.Sprintf("%d\n", os.Getpid())), 0666)
}

// Remove 删除PID文件
func Remove(filename string) error {
	return os.Remove(filename)
}
