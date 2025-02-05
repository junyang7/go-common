package _mod

import (
	"fmt"
	"github.com/junyang7/go-common/src/_cmd"
	"github.com/junyang7/go-common/src/_file"
	"github.com/junyang7/go-common/src/_interceptor"
	"golang.org/x/mod/modfile"
	"os"
)

func Init() {

	path := ""
	if _, ok := os.LookupEnv("GO_TEST"); ok {
		path = "../../go.mod"
	} else {
		path = "go.mod"
	}
	_interceptor.Insure(_file.Exists(path)).Message("go.mod文件不存在").Do()

	f, err := modfile.Parse(path, _file.ReadAll(path), nil)
	_interceptor.Insure(nil == err).Message(err).Do()

	for _, req := range f.Require {
		if !req.Indirect {
			continue
		}
		cmd := fmt.Sprintf(`go get %s@%s`, req.Mod.Path, req.Mod.Version)
		fmt.Println(cmd)
		_cmd.ExecuteByStd(cmd)
	}

}
