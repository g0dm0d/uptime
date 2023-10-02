package uptime

import (
	"github.com/g0dm0d/uptime/internal/server/socket"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/model"
	"github.com/g0dm0d/uptime/pkg/cron"
)

type Uptime struct {
	cron           cron.Cron
	monitorStore   store.MonitorStore
	heartbeatStore store.HeartbeatStore
	websocket      *socket.Socket
}

func New(cron cron.Cron, hs store.HeartbeatStore, ms store.MonitorStore, ws *socket.Socket) *Uptime {
	return &Uptime{
		cron:           cron,
		heartbeatStore: hs,
		monitorStore:   ms,
		websocket:      ws,
	}
}

func (u *Uptime) Init() error {
	monitors, err := u.monitorStore.GetAllMonitor()
	if err != nil {
		return err
	}
	for _, monitor := range monitors {
		u.cron.AddTask(cron.Task{
			MonitorID: monitor.ID,
			Schedule: cron.Schedule{
				IsDate:  false,
				Day:     0,
				Hours:   0,
				Minutes: 0,
				Seconds: monitor.Interval,
			},
			Action: u.getPingFunc(monitor.Protocol, monitor.ID),
		})
	}
	u.cron.Start()
	return nil
}

func (u *Uptime) AddMonitor(monitor *model.Monitor) error {
	task := cron.Task{
		MonitorID: monitor.ID,
		Schedule: cron.Schedule{
			IsDate:  false,
			Day:     0,
			Hours:   0,
			Minutes: 0,
			Seconds: monitor.Interval,
		},
		Action: u.getPingFunc(store.Protocol(monitor.Protocol), monitor.ID),
	}
	u.cron.AddTask(task)
	err := u.cron.RunByID(task.MonitorID)
	return err
}

func (u *Uptime) getPingFunc(protocol store.Protocol, monitorID int) func(int) error {
	switch protocol {
	case store.TCP:
		return u.PingTCP
	}
	return u.PingTCP
}
