package model

type Heartbeat struct {
	Alive     bool
	Message   string
	FreeSpace int64
}
