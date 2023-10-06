package store

import "time"

type Tick struct {
	MonitorID string `json:"monitor_id"`
	Success   int    `json:"success"`
	Ping      int    `json:"ping"`
	Msg       string `json:"msg"`
}

type Heartbeat struct {
	Date time.Time
	Tick Tick
}

type HeartbeatStore interface {
	SaveTick(Tick) error
	GetTickHistory(monitorID string, count, timestamp int) ([]Heartbeat, error)
}
