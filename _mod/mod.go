package _mod

import (
	"fmt"
	"github.com/junyang7/go-common/_cmd"
	"github.com/junyang7/go-common/_file"
	"github.com/junyang7/go-common/_interceptor"
	"golang.org/x/mod/modfile"
)

func Init(path string) {

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
