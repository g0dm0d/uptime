package model

import "github.com/g0dm0d/uptime/internal/store"

type Monitor struct {
	ID       int
	Hostname string
	Interval int
	Protocol string
	Addr     string
	Port     interface{}
	Tags     []string
}

func NewMonitor(m store.Monitor) Monitor {
	monitor := Monitor{
		ID:       m.ID,
		Hostname: m.Hostname,
		Protocol: string(m.Protocol),
		Addr:     m.Addr,
		Port:     m.Port.Int16,
		Tags:     m.Tags,
	}

	if !m.Port.Valid {
		monitor.Port = nil
	}
	return monitor
}

func NewMonitors(m []store.Monitor) []Monitor {
	var monitors []Monitor
	for i := range m {
		monitors = append(monitors, NewMonitor(m[i]))
	}
	return monitors
}
