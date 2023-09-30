package store

type Heartbeat struct {
	MonitorID int    `json:"monitor_id"`
	Success   int    `json:"success"`
	Ping      int    `json:"ping"`
	Msg       string `json:"msg"`
}

type HeartbeatStore interface {
	SaveTick(Heartbeat) error
	GetTickHistory(monitorID, count int) ([]Heartbeat, error)
}
