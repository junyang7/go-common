package _port

import (
	"fmt"
	"os/exec"
)

func GetList(count int) []int {
	var portList []int
	var i int
	for i = 10000; i < 65535; i++ {
		handler := exec.Command(`lsof`, fmt.Sprintf(`-i:%d`, i))
		_, err := handler.Output()
		if err != nil {
			portList = append(portList, i)
			if len(portList) >= count {
				break
			}
		}
	}
	return portList
}
