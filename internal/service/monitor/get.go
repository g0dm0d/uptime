package monitor

import (
	"github.com/g0dm0d/uptime/dto"
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/go-chi/chi/v5"
)

func (s *Service) GetAll(ctx *req.Ctx) error {
	monitors := s.uptime.GetMonitorList()
	return ctx.JSON(dto.NewMonitors(monitors))
}

func (s *Service) Get(ctx *req.Ctx) error {
	monitorID := chi.URLParam(ctx.Request, "monitor")

	monitor := s.uptime.GetMonitor(monitorID)
	return ctx.JSON(dto.NewMonitor(monitor))
}
