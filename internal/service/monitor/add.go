package monitor

// import (
// 	"log"
//
// 	"github.com/g0dm0d/uptime/internal/server/req"
// 	"github.com/g0dm0d/uptime/internal/store"
// 	"github.com/g0dm0d/uptime/model"
// 	"github.com/g0dm0d/uptime/pkg/errs"
// )
//
// type AddMonitorReq struct {
// 	Hostname string         `json:"hostname"`
// 	Interval int            `json:"interval"`
// 	Protocol store.Protocol `json:"protocol"`
// 	Addr     string         `json:"address"`
// 	Port     interface{}    `json:"port"`
// 	Tags     []string       `json:"tags"`
// }
//
// func (s *Service) Add(ctx *req.Ctx) error {
// 	var r AddMonitorReq
//
// 	err := ctx.ParseJSON(&r)
// 	if err != nil {
// 		return errs.ReturnError(ctx.Writer, errs.InvalidJSON)
// 	}
//
// 	if ok := r.Protocol.IsValid(); !ok {
// 		return errs.ReturnError(ctx.Writer, errs.InvalidProtocol)
// 	}
//
// 	id, err := s.monitorStore.AddMonitor(store.CreateMonitorOpts{
// 		Hostname: r.Hostname,
// 		Interval: r.Interval,
// 		Protocol: string(r.Protocol),
// 		Addr:     r.Addr,
// 		Port:     r.Port,
// 		Tags:     r.Tags,
// 	})
//
// 	if err != nil {
// 		log.Println(err)
// 		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
// 	}
//
// 	err = s.uptime.AddMonitor(&model.Monitor{
// 		ID:       id,
// 		Hostname: r.Hostname,
// 		Interval: r.Interval,
// 		Protocol: string(r.Protocol),
// 		Addr:     r.Addr,
// 		Port:     r.Port,
// 		Tags:     r.Tags,
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return errs.ReturnError(ctx.Writer, errs.InternalServerError)
// 	}
//
// 	return ctx.JSON(nil)
// }
