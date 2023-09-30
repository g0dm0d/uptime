package uptime

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/g0dm0d/uptime/internal/store"
)

func (u *Uptime) PingTCP(monitorID int) error {
	start := time.Now()

	monitor, err := u.monitorStore.GetMonitor(monitorID)
	if err != nil {
		return err
	}

	port, _ := monitor.Port.Value()
	conn, err := net.DialTimeout("tcp", fmt.Sprint(monitor.Addr, ":", port), 1*time.Second)
	if err != nil {
		result := store.Heartbeat{
			MonitorID: monitorID,
			Success:   0,
			Ping:      0,
			Msg:       err.Error(),
		}

		resultJSON, err := json.Marshal(result)
		if err != nil {
			return err
		}

		u.websocket.Emit(string(resultJSON))

		err = u.heartbeatStore.SaveTick(result)
		return err
	}
	defer conn.Close()

	elapsedTime := time.Since(start)

	result := store.Heartbeat{
		MonitorID: monitorID,
		Success:   1,
		Ping:      int(elapsedTime.Milliseconds()),
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	u.websocket.Emit(string(resultJSON))

	err = u.heartbeatStore.SaveTick(result)
	return err
}
