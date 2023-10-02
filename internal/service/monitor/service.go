package monitor

import (
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/internal/uptime"
)

type Monitor interface {
	Add(ctx *req.Ctx) error
	GetAll(ctx *req.Ctx) error
	GetHistory(ctx *req.Ctx) error
}

type Service struct {
	monitorStore   store.MonitorStore
	heartbeatStore store.HeartbeatStore
	uptime         uptime.Uptime
}

func New(ms store.MonitorStore, hs store.HeartbeatStore, u uptime.Uptime) *Service {
	return &Service{
		monitorStore:   ms,
		heartbeatStore: hs,
		uptime:         u,
	}
}
