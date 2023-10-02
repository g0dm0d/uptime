package uptime

import (
	"fmt"
	"net"
	"time"
)

func (u *Uptime) PingTCP(monitorID string) error {
	start := time.Now()

	monitor, ok := u.monitors[monitorID]
	if !ok {
		return fmt.Errorf("id not found")
	}

	conn, err := net.DialTimeout("tcp", fmt.Sprint(monitor.Addr, ":", monitor.Port), time.Duration(monitor.Interval)*time.Second)
	if err != nil {
		tick := GenerateFailTick(monitorID, err.Error())
		return u.SaveAndEmitTick(tick)
	}
	defer conn.Close()

	elapsedTime := time.Since(start)

	tick := GenerateSuccessTick(monitorID, "", int(elapsedTime.Milliseconds()))

	return u.SaveAndEmitTick(tick)
}
