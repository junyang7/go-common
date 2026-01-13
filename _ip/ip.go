package _ip

import (
	"github.com/junyang7/go-common/_interceptor"
	"net"
)

func GetByLocal() string {
	return GetListByLocal()[0]
}
func GetListByLocal() []string {
	var res []string
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		_interceptor.Insure(false).Message(err).Do()
		return res
	}
	for _, addr := range addrList {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipNet.IP.To4() != nil && !ipNet.IP.IsLoopback() {
			res = append(res, ipNet.IP.String())
		}
	}
	if len(res) == 0 {
		res = append(res, "127.0.0.1")
	}
	return res
}
