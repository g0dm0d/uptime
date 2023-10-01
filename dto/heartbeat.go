package dto

import (
	"time"

	"github.com/g0dm0d/uptime/model"
)

type Heartbeat struct {
	MonitorID int       `json:"monitor_id"`
	Success   int       `json:"success"`
	Ping      int       `json:"ping"`
	Msg       string    `json:"message"`
	Date      time.Time `json:"date"`
}

func NewHeartbeat(h model.Heartbeat) Heartbeat {
	return Heartbeat{
		MonitorID: h.MonitorID,
		Success:   h.Success,
		Ping:      h.Success,
		Msg:       h.Msg,
		Date:      h.Date,
	}
}

func NewHeartbeats(h []model.Heartbeat) []Heartbeat {
	var heartbeats []Heartbeat
	for i := range h {
		heartbeats = append(heartbeats, NewHeartbeat(h[i]))
	}
	return heartbeats
}
