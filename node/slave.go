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
	Strategy Strategy `json:"strategy"`
	Interval string   `json:"interval"`
}

type Strategy string

const (
	Periodic Strategy = "periodic"
)

func (c SlaveConfig) GetInterval() time.Duration {
	d, err := time.ParseDuration(c.Interval)
	if err != nil {
		panic(err)
	}
	return d

}

func (c SlaveConfig) Run() {
	dir, _ := os.Getwd()
	b := backend.GitBackend{
		Base: backend.Base{Location: dir},
	}
	dur := c.GetInterval()
	fmt.Println("Pulling to head")
	b.PullToHead()
	for {
		select {
		case <-time.After(dur):
			fmt.Println("Pushing")
			b.Push(false)
		}
	}

}
