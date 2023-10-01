package uptime

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/g0dm0d/uptime/dto"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/model"
)

func (u *Uptime) PingTCP(monitorID int) error {
	start := time.Now()

	monitor, err := u.monitorStore.GetMonitor(monitorID)
	if err != nil {
		return err
	}

	port, _ := monitor.Port.Value()
	conn, err := net.DialTimeout("tcp", fmt.Sprint(monitor.Addr, ":", port), time.Duration(monitor.Interval)*time.Second)
	if err != nil {
		tick := store.Tick{
			MonitorID: monitorID,
			Success:   0,
			Ping:      0,
			Msg:       err.Error(),
		}

		result := dto.NewHeartbeat(model.NewHeartbeat(store.Heartbeat{
			Date: start,
			Tick: tick,
		}))

		resultJSON, err := json.Marshal(result)
		if err != nil {
			return err
		}

		u.websocket.Emit(string(resultJSON))

		err = u.heartbeatStore.SaveTick(tick)
		return err
	}
	defer conn.Close()

	elapsedTime := time.Since(start)

	tick := store.Tick{
		MonitorID: monitorID,
		Success:   1,
		Ping:      int(elapsedTime.Milliseconds()),
	}

	result := dto.NewHeartbeat(model.NewHeartbeat(store.Heartbeat{
		Date: start,
		Tick: tick,
	}))

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	u.websocket.Emit(string(resultJSON))

	err = u.heartbeatStore.SaveTick(tick)
	return err
}
