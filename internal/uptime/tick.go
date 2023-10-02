package uptime

import (
	"encoding/json"
	"time"

	"github.com/g0dm0d/uptime/dto"
	"github.com/g0dm0d/uptime/internal/store"
	"github.com/g0dm0d/uptime/model"
)

func GenerateFailTick(monitorID string, msg string) store.Tick {
	return store.Tick{
		MonitorID: monitorID,
		Success:   0,
		Ping:      0,
		Msg:       msg,
	}
}

func GenerateSuccessTick(monitorID string, msg string, ping int) store.Tick {
	return store.Tick{
		MonitorID: monitorID,
		Success:   1,
		Ping:      ping,
		Msg:       msg,
	}
}

func (u *Uptime) SaveAndEmitTick(tick store.Tick) error {
	result := dto.NewHeartbeat(model.NewHeartbeat(store.Heartbeat{
		Date: time.Now(),
		Tick: tick,
	}))

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return err
	}

	u.websocket.Emit(string(resultJSON))

	return u.heartbeatStore.SaveTick(tick)
}
