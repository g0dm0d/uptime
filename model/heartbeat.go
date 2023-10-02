package model

import (
	"time"

	"github.com/g0dm0d/uptime/internal/store"
)

type Heartbeat struct {
	MonitorID string
	Success   int
	Ping      int
	Msg       string
	Date      time.Time
}

func NewHeartbeat(h store.Heartbeat) Heartbeat {
	return Heartbeat{
		MonitorID: h.Tick.MonitorID,
		Success:   h.Tick.Success,
		Ping:      h.Tick.Ping,
		Msg:       h.Tick.Msg,
		Date:      h.Date,
	}
}

func NewHeartbeats(h []store.Heartbeat) []Heartbeat {
	var heartbeats []Heartbeat
	for i := range h {
		heartbeats = append(heartbeats, NewHeartbeat(h[i]))
	}
	return heartbeats
}
