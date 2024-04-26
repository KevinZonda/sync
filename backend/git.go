package backend

import (
	"fmt"
	"github.com/KevinZonda/sync/common"
	"strings"
)

type GitBackend struct {
	Base
}

func (g *GitBackend) PullToHead() error {
	return g.runCmd("git", "pull")
}

func (g *GitBackend) Push(force bool) error {
	if !g.needAddOrCommit() {
		fmt.Println("Nothing to do")
		return nil
	}
	g.runCmd("git", "add", ".")
	g.runCmd("git", "commit", "-m", common.NamingCommit())
	if !g.needPush() {
		return nil
	}
	fmt.Println("Need push")
	if force {
		return g.runCmd("git", "push", "--force")
	}
	return g.runCmd("git", "push")
}

func (g *GitBackend) needAddOrCommit() bool {
	content, _ := g.cmdStr("git", "status", "--porcelain")
	return content != ""
}

func (g *GitBackend) needPush() bool {
	local, _ := g.cmdStr("git", "rev-parse", "HEAD")
	remote, _ := g.cmdStr("git", "rev-parse", "@{u}")
	local = strings.TrimSpace(local)
	remote = strings.TrimSpace(remote)
	fmt.Println("Checking git hash: ", local, remote)
	return local != remote
}

var _ common.IBaseSyncBackend = (*GitBackend)(nil)
