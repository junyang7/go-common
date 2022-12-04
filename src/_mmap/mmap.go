package _mmap

import (
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
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
		panic(err)
	}
	data, err := syscall.Mmap(int(file.Fd()), offset, size, prot, flag)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSyscallMmap).
		Message(err).
		Do()
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
	_, _, err := syscall.Syscall(syscall.SYS_MSYNC, uintptr(unsafe.Pointer(&this.data[0])), uintptr(this.offset), uintptr(ms))
	_interceptor.Insure(0 == err).
		CodeMessage(_codeMessage.ErrSyscallSyscall).
		Message(err).
		Do()
}
func (this *Mmap) Close() {
	err := syscall.Munmap(this.data)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrSyscallMunmap).
		Message(err).
		Do()
	this.data = nil
}
