package monitor

import (
	"github.com/g0dm0d/uptime/dto"
	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/model"
	"github.com/g0dm0d/uptime/pkg/errs"
	"github.com/go-chi/chi/v5"
)

type HistoryResp struct {
}

func (s *Service) GetHistory(ctx *req.Ctx) error {
	monitor := chi.URLParam(ctx.Request, "monitor")

	data, err := s.heartbeatStore.GetTickHistory(monitor, 20)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}
	return ctx.JSON(dto.NewHeartbeats(model.NewHeartbeats(data)))
}
