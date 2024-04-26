package backend

import "github.com/KevinZonda/sync/common"

type GitBackend struct {
	Base
}

func (g *GitBackend) PullToHead() error {
	return g.runCmd("git", "pull")
}

func (g *GitBackend) Push(force bool) error {
	if force {
		return g.runCmd("git", "push", "--force")
	}
	return g.runCmd("git", "push")
}

var _ common.IBaseSyncBackend = (*GitBackend)(nil)
