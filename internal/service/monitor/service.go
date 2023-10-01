package monitor

import (
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/internal/store"
)

type Monitor interface {
	Add(ctx *req.Ctx) error
	GetAll(ctx *req.Ctx) error
	GetHistory(ctx *req.Ctx) error
}

type Service struct {
	monitorStore   store.MonitorStore
	heartbeatStore store.HeartbeatStore
}

func New(monitorStore store.MonitorStore, heartbeatStore store.HeartbeatStore) *Service {
	return &Service{
		monitorStore:   monitorStore,
		heartbeatStore: heartbeatStore,
	}
}
