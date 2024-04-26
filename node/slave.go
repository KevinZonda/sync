package node

import (
	"fmt"
	"github.com/KevinZonda/sync/backend"
	"os"
	"time"
)

type Slave struct {
}

func (s *Slave) ReportToMaster() {

}

type SlaveConfig struct {
	Strategy Strategy
	Minutes  time.Duration
}

type Strategy string

const (
	Periodic Strategy = "periodic"
)

func (c SlaveConfig) Run() {
	dir, _ := os.Getwd()
	b := backend.GitBackend{
		Base: backend.Base{Location: dir},
	}
	fmt.Println("Pulling to head")
	b.PullToHead()
	for {
		select {
		case <-time.After(c.Minutes):
			fmt.Println("Pushing")
			b.Push(false)
		}
	}

}
