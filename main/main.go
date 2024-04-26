package main

import (
	"fmt"
	"github.com/KevinZonda/GoX/pkg/iox"
	"github.com/KevinZonda/sync/node"
	"github.com/KevinZonda/sync/text"
	"os"
	"time"
)

func main() {
	txt := ""
	var err error
	for _, v := range text.SYNC_SLAVE_FILE {
		txt, err = iox.ReadAllText(v)
		if err != nil || txt == "" {
			continue
		}
	}
	if txt == "" {
		fmt.Println("no valid config file")
		os.Exit(1)
	}
	c := node.SlaveConfig{
		Strategy: node.Periodic,
		Minutes:  10 * time.Second,
	}
	c.Run()
}
