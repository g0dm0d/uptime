package monitor

import (
	"github.com/g0dm0d/uptime/dto"
	"github.com/g0dm0d/uptime/internal/server/req"
)

func (s *Service) GetAll(ctx *req.Ctx) error {
	monitors := s.uptime.GetMonitorList()
	return ctx.JSON(dto.NewMonitors(monitors))
}
