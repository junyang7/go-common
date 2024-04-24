package _ip

import "net"

func GetByLocal() string {
	res := "127.0.0.1"
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrList {
		ipNet, _ := addr.(*net.IPNet)
		if ipNet.IP.To4() != nil && !ipNet.IP.IsLoopback() {
			res = ipNet.IP.String()
		}
	}
	return res
}
