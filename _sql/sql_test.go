package _sql

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	{
		machine := &Machine{
			Driver:    `mysql`,
			Host:      `127.0.0.1`,
			Port:      `3306`,
			Database:  `me`,
			Username:  `root`,
			Password:  ``,
			Charset:   ``,
			Collation: ``,
		}
		res := New().Machine(machine).Sql("select * from link limit 1").Query()
		b, _ := json.MarshalIndent(res, " ", " ")
		fmt.Println(string(b))
	}
	{
		machine := &Machine{
			Driver:    `sqlite3`,
			Host:      ``,
			Port:      ``,
			Database:  `/Users/junyang7/env/env.db`,
			Username:  ``,
			Password:  ``,
			Charset:   ``,
			Collation: ``,
		}
		res := New().Machine(machine).Table("project").Get()
		b, _ := json.MarshalIndent(res, " ", " ")
		fmt.Println(string(b))
	}
}
