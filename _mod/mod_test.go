package _mod

import (
	"github.com/junyang7/go-common/_directory"
	"testing"
)

func TestInit(t *testing.T) {
	Init(_directory.Name(_directory.Current(), 1) + "/go.mod")
}
