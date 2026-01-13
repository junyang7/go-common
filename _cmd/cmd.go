package _cmd

import (
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_interceptor"
	"os"
	"os/exec"
)

func Execute(name string, arg ...string) []byte {
	cmd := exec.Command(name, arg...)
	b, err := cmd.Output()
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return b
}
func ExecuteAsInt64(name string, arg ...string) int64 {
	return _as.Int64(Execute(name, arg...))
}
func ExecuteAsString(name string, arg ...string) string {
	return _as.String(Execute(name, arg...))
}
func ExecuteByStd(cmd string) {
	handler := exec.Command("/bin/bash", "-c", cmd)
	handler.Stdout = os.Stdout
	handler.Stderr = os.Stderr
	if err := handler.Run(); err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
}
