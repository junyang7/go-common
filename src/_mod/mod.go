package _mod

import (
	"fmt"
	"github.com/junyang7/go-common/src/_file"
	"path/filepath"
)

func Init() {
	path := "../../go.mod"
	fmt.Println(filepath.Abs(path))
	if _file.Exists(path) {
		fmt.Println("ok")
	} else {
		fmt.Println("go.mod文件不存在")
	}
}
