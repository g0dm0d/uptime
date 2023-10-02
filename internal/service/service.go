package service

import (
	"github.com/g0dm0d/uptime/internal/service/monitor"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/internal/uptime"
)

type Service struct {
	Monitor monitor.Monitor
}

type Opts struct {
	HeartbeatStore store.HeartbeatStore
	Uptime         uptime.Uptime
}

func New(s Opts) *Service {
	return &Service{
		Monitor: monitor.New(s.HeartbeatStore, s.Uptime),
	}
}
