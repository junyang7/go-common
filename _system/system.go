package _system

import (
	"github.com/junyang7/go-common/_interceptor"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func Memory() *mem.VirtualMemoryStat {

	stat, err := mem.VirtualMemory()
	_interceptor.Insure(nil == err).Message(err).Do()
	return stat

}

func Disk(path string) *disk.UsageStat {

	stat, err := disk.Usage(path)
	_interceptor.Insure(nil == err).Message(err).Do()
	return stat

}
