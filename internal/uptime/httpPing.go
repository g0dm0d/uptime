package uptime

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (u *Uptime) PingHTTP(monitorID string) error {
	start := time.Now()

	monitor, ok := u.monitors[monitorID]
	if !ok {
		return fmt.Errorf("id not found")
	}

	var reqBodyReader io.Reader = nil

	if monitor.Body != "" {
		reqBodyReader = bytes.NewReader([]byte(monitor.Body))
	}

	req, err := http.NewRequest(monitor.Method, fmt.Sprintf("%s://%s", monitor.Protocol, monitor.Addr), reqBodyReader)
	if err != nil {
		tick := GenerateFailTick(monitorID, err.Error())
		return u.SaveAndEmitTick(tick)
	}

	for _, header := range monitor.Headers {
		req.Header.Add(header[0], header[1])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		tick := GenerateFailTick(monitorID, err.Error())
		return u.SaveAndEmitTick(tick)
	}
	defer resp.Body.Close()

	// _, err = io.ReadAll(resp.Body)
	// if err != nil {
	// 	return err
	// }

	if resp.StatusCode != monitor.Status {
		tick := GenerateFailTick(monitorID, resp.Status)
		return u.SaveAndEmitTick(tick)
	}

	elapsedTime := time.Since(start)

	tick := GenerateSuccessTick(monitorID, resp.Status, int(elapsedTime.Milliseconds()))

	return u.SaveAndEmitTick(tick)
}
