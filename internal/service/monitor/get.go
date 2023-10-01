package monitor

import (
	"github.com/g0dm0d/uptime/dto"
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/model"
	"github.com/g0dm0d/uptime/pkg/errs"
)

func (s *Service) GetAll(ctx *req.Ctx) error {
	monitors, err := s.monitorStore.GetAllMonitor()
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}

	return ctx.JSON(dto.NewMonitors(model.NewMonitors(monitors)))
}
