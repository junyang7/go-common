package _unixMilli

import "time"

func Get() int64 {
	return time.Now().UnixMicro()
}
