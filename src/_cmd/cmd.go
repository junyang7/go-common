package _cmd

import (
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"os/exec"
)

func Execute(name string, arg ...string) []byte {
	cmd := exec.Command(name, arg...)
	b, err := cmd.Output()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrCmdOutput).
		Message(err.Error()).
		Do()
	return b
}
func ExecuteAsInt64(name string, arg ...string) int64 {
	return _as.Int64(Execute(name, arg...))
}
func ExecuteAsString(name string, arg ...string) string {
	return _as.String(Execute(name, arg...))
}
