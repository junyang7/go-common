package _conf

import "database/sql"

type Machine struct {
	Driver    string `json:"driver"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	Database  string `json:"database"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	Collation string `json:"collation"`
	Pool      *sql.DB
}
