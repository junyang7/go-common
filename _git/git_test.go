package _git

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestIsBranchExistsInRemote(t *testing.T) {

	{
		repository := "/Users/junyang7/env/27-发票管理/home/q/system/cmn_order"
		branch := "master"
		var expect bool = true
		get := IsBranchExistsInRemote(repository, branch)
		_assert.Equal(t, expect, get)
	}
	{
		repository := "/Users/junyang7/env/27-发票管理/home/q/system/cmn_order"
		branch := "master1"
		var expect bool = false
		get := IsBranchExistsInRemote(repository, branch)
		_assert.Equal(t, expect, get)
	}

}
