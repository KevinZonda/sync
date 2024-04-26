package common

type IBaseSyncBackend interface {
	PullToHead() error
	Push(force bool) error
}
