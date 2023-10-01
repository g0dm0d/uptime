package dto

import "github.com/g0dm0d/uptime/model"

type Monitor struct {
	ID       int         `json:"id"`
	Hostname string      `json:"hostname"`
	Interval int         `json:"interval"`
	Protocol string      `json:"protocol"`
	Addr     string      `json:"address"`
	Port     interface{} `json:"port"`
	Tags     []string    `json:"tags"`
}

func NewMonitor(m model.Monitor) Monitor {
	return Monitor{
		ID:       m.ID,
		Hostname: m.Hostname,
		Protocol: string(m.Protocol),
		Addr:     m.Addr,
		Port:     m.Port,
		Tags:     m.Tags,
	}
}

func NewMonitors(m []model.Monitor) []Monitor {
	var users []Monitor
	for i := range m {
		users = append(users, NewMonitor(m[i]))
	}
	return users
}
