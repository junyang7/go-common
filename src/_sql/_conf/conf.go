package _conf

var (
	Conf *struct {
		Database map[string]*struct {
			Count   int `json:"count"`
			Cluster map[string]*struct {
				Master *struct {
					Count   int        `json:"count"`
					Machine []*Machine `json:"machine"`
				} `json:"master"`
				Slaver *struct {
					Count   int        `json:"count"`
					Machine []*Machine `json:"machine"`
				} `json:"slaver"`
			} `json:"cluster"`
		} `json:"database"`
		Table map[string]int `json:"table"`
	}
)
