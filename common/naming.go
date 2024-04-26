package common

import "time"

func NamingCommit() string {
	return time.Now().String()
}
