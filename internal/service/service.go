package service

import (
	"github.com/g0dm0d/uptime/internal/service/monitor"
	"github.com/g0dm0d/uptime/internal/service/user"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/internal/uptime"
	"github.com/g0dm0d/uptime/pkg/jwtmanager"
)

type Service struct {
	User    user.User
	Monitor monitor.Monitor
}

type Opts struct {
	UserStore      store.UserStore
	MonitorStore   store.MonitorStore
	HeartbeatStore store.HeartbeatStore
	Uptime         uptime.Uptime
	JWT            jwtmanager.Tool
}

func New(s Opts) *Service {
	return &Service{
		User:    user.New(s.UserStore, s.JWT),
		Monitor: monitor.New(s.MonitorStore, s.HeartbeatStore, s.Uptime),
	}
}
