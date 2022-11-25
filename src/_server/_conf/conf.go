package _conf

type Conf struct {
	Ip   string
	Port string
	Ipv4 struct {
		Black []string
		White []string
	}
	Method struct {
		Black []string
		White []string
	}
	Origin []string
	Header map[string]string
}
