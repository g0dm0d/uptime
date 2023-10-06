package monitor

import (
	"strconv"

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
	countStr := ctx.Request.URL.Query().Get("count")
	timeFromStr := ctx.Request.URL.Query().Get("from")

	count, err := strconv.Atoi(countStr)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidUrlParam)
	}

	TimeFrom, err := strconv.Atoi(timeFromStr)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InvalidUrlParam)
	}

	data, err := s.heartbeatStore.GetTickHistory(monitor, count, TimeFrom)
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
	}
	return ctx.JSON(dto.NewHeartbeats(model.NewHeartbeats(data)))
}
