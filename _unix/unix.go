package _unix

import "time"

func Get() int64 {
	return time.Now().Unix()
}
