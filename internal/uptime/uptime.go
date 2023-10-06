package uptime

import (
	"github.com/g0dm0d/uptime/internal/server/socket"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/model"
	"github.com/g0dm0d/uptime/pkg/cron"
)

type Protocol string

const (
	HTTP Protocol = "http"
	TCP  Protocol = "tcp"
)

func (p Protocol) IsValid() bool {
	switch p {
	case HTTP, TCP:
		return true
	}
	return false
}

type Uptime struct {
	cron           cron.Cron
	monitors       map[string]store.MonitorConfig
	heartbeatStore store.HeartbeatStore
	websocket      *socket.Socket
}

func New(cron cron.Cron, hs store.HeartbeatStore, ws *socket.Socket) *Uptime {
	return &Uptime{
		cron:           cron,
		heartbeatStore: hs,
		monitors:       make(map[string]store.MonitorConfig),
		websocket:      ws,
	}
}

func (u *Uptime) Init(config store.UptimeConfig) error {
	for _, server := range config.GetList() {
		u.load(server)
		u.monitors[server.ID] = server
	}
	u.cron.Start()
	return nil
}

func (u *Uptime) GetMonitorList() []model.Monitor {
	monitors := []model.Monitor{}
	for _, monitor := range u.monitors {
		monitors = append(monitors, model.NewMonitor(monitor))
	}
	return monitors
}

func (u *Uptime) GetMonitor(monitorID string) model.Monitor {
	if monitor, ok := u.monitors[monitorID]; ok {
		return model.NewMonitor(monitor)
	}
	return model.Monitor{}
}

func (u *Uptime) load(server store.MonitorConfig) {
	u.cron.AddTask(cron.Task{
		MonitorID: server.ID,
		Schedule: cron.Schedule{
			IsDate:  false,
			Day:     0,
			Hours:   0,
			Minutes: 0,
			Seconds: server.Interval,
		},
		Action: u.getPingFunc(server.Protocol),
	})
}

func (u *Uptime) getPingFunc(protocol store.Protocol) func(string) error {
	switch protocol {
	case store.TCP:
		return u.PingTCP
	case store.HTTP:
		return u.PingHTTP
	case store.HTTPS:
		return u.PingHTTP
	}
	return u.PingTCP
}
