package _mmap

import (
	"github.com/junyang7/go-common/_interceptor"
	"os"
	"syscall"
	"unsafe"
)

type Mmap struct {
	file   *os.File
	offset int64
	data   []byte
}

func Open(file *os.File, offset int64, size int, prot int, flag int) *Mmap {
	if err := file.Truncate(int64(size)); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	data, err := syscall.Mmap(int(file.Fd()), offset, size, prot, flag)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return &Mmap{file: file, data: data}
}
func (this *Mmap) Read(offset int64, length int64) []byte {
	return this.data[offset:length]
}
func (this *Mmap) Append(b []byte) {
	for _, v := range b {
		this.data[this.offset] = v
		this.offset++
	}
}
func (this *Mmap) Flush(ms int) {
	if _, _, err := syscall.Syscall(syscall.SYS_MSYNC, uintptr(unsafe.Pointer(&this.data[0])), uintptr(this.offset), uintptr(ms)); 0 != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
func (this *Mmap) Close() {
	err := syscall.Munmap(this.data)
	if err != nil {
		_interceptor.Insure(false).Message(err).Do()
	}
	this.data = nil
}
