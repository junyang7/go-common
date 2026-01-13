package _log

import "testing"

func TestDebug(t *testing.T) {
	Debug("a", true, 1)
}
func TestInfo(t *testing.T) {
	Info("a", true, 1)
}
func TestWarning(t *testing.T) {
	Warning("a", true, 1)
}
func TestError(t *testing.T) {
	Error("a", true, 1)
}
func TestWrite(t *testing.T) {
	Write("custom", "a", true, 1)
}
