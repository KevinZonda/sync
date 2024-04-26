package text

import (
	"os"
	"path"
)

const IDENTIFIER = "com.kevinzonda.sync"
const NAME = "sync"
const AUTHOR = "KevinZonda"
const VERSION = "0.0.0"

var SYNC_SLAVE_FILE = []string{"config.ksync", ".ksync"}

func ConfigPath(file ...string) string {
	d, _ := os.UserConfigDir()
	ds := []string{d, IDENTIFIER}
	if len(file) == 0 {
		return path.Join(ds...)
	}
	return path.Join(append(ds, file...)...)
}

func Info() string {
	return NAME + " - " + VERSION + " by " + AUTHOR
}
