package _millisecond

import "time"

func Get() int64 {
	return time.Now().UnixNano() / 1e6
}
