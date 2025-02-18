package _port

import (
	"fmt"
	"github.com/junyang7/go-common/_list"
	"os/exec"
)

func IsUsing(port int) bool {
	handler := exec.Command(`lsof`, fmt.Sprintf(`-i:%d`, port))
	_, err := handler.Output()
	return err != nil
}
func GetList(count int) (o []int) {
	for i := 1; i <= 65535; i++ {
		if !IsUsing(i) {
			o = append(o, i)
			if len(o) >= count {
				break
			}
		}
	}
	return o
}
func Get() (o int) {
	return GetList(1)[0]
}

type port struct {
	min    int
	max    int
	count  int
	filter []int
}

func New() *port {
	return &port{
		min:    20000,
		max:    65535,
		count:  10,
		filter: []int{},
	}
}
func (this *port) Min(min int) *port {
	this.min = min
	return this
}
func (this *port) Max(max int) *port {
	this.max = max
	return this
}
func (this *port) Count(count int) *port {
	this.count = count
	return this
}
func (this *port) Filter(filter []int) *port {
	this.filter = append(this.filter, filter...)
	return this
}
func (this *port) GetList() (o []int) {
	for i := this.min; i <= this.max; i++ {
		if _list.In(i, this.filter) {
			continue
		}
		if !IsUsing(i) {
			o = append(o, i)
			if len(o) >= this.count {
				break
			}
		}
	}
	return o
}
