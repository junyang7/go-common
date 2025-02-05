package _mod

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	os.Setenv("GO_TEST", "true")
	Init()
}
