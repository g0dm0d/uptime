package monitor

import (
	"strconv"

	"github.com/g0dm0d/uptime/internal/server/req"
	"github.com/g0dm0d/uptime/pkg/errs"
	"github.com/go-chi/chi/v5"
)

type HistoryResp struct {
}

func (s *Service) GetHistory(ctx *req.Ctx) error {
	monitor := chi.URLParam(ctx.Request, "monitor")
	i, err := strconv.Atoi(monitor)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidUrlParam)
	}
	data, err := s.heartbeatStore.GetTickHistory(i, 20)
	return ctx.JSON(data)
}
