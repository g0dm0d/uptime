package model

import "github.com/g0dm0d/uptime/internal/store"

type Monitor struct {
	ID       string
	Hostname string
	Interval int
	Protocol string
	Addr     string
	Port     interface{}
}

func NewMonitor(s store.MonitorConfig) Monitor {
	monitor := Monitor{
		ID:       s.ID,
		Hostname: s.Name,
		Protocol: string(s.Protocol),
		Addr:     s.Addr,
		Port:     s.Port,
	}

	if s.Port < 1 {
		monitor.Port = nil
	}
	return monitor
}

func NewMonitors(s []store.MonitorConfig) []Monitor {
	var monitors []Monitor
	for i := range s {
		monitors = append(monitors, NewMonitor(s[i]))
	}
	return monitors
}
