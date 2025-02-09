package _structure

type ServerWeb struct {
	Debug  bool     `json:"debug"`
	Host   string   `json:"host"`
	Port   string   `json:"port"`
	Origin []string `json:"origin"`
	Root   string   `json:"root"`
}

type ServerApi struct {
	Debug  bool     `json:"debug"`
	Host   string   `json:"host"`
	Port   string   `json:"port"`
	Origin []string `json:"origin"`
}

type ServerHttp struct {
	Debug  bool     `json:"debug"`
	Host   string   `json:"host"`
	Port   string   `json:"port"`
	Origin []string `json:"origin"`
	Root   string   `json:"root"`
}
