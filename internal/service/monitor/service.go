package monitor

import (
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/internal/uptime"
)

type Monitor interface {
	Get(ctx *req.Ctx) error
	GetAll(ctx *req.Ctx) error
	GetHistory(ctx *req.Ctx) error
}

type Service struct {
	heartbeatStore store.HeartbeatStore
	uptime         uptime.Uptime
}

func New(hs store.HeartbeatStore, u uptime.Uptime) *Service {
	return &Service{
		heartbeatStore: hs,
		uptime:         u,
	}
}
