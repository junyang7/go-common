package _ip

import (
	"github.com/junyang7/go-common/src/_interceptor"
	"net"
)

func GetByLocal() string {
	res := "127.0.0.1"
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
	for _, addr := range addrList {
		ipNet, _ := addr.(*net.IPNet)
		if ipNet.IP.To4() != nil && !ipNet.IP.IsLoopback() {
			res = ipNet.IP.String()
			break
		}
	}
	return res
}
