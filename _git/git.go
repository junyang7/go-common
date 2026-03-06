package _git

import (
	"github.com/junyang7/go-common/_cmd"
)

func IsBranchExistsInRemote(repository string, branch string) bool {
	b := _cmd.Execute("git", "-C", repository, "ls-remote", "--heads", "origin", branch)
	return len(b) > 0
}
