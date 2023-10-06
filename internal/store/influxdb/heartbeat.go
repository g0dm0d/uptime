package influxdb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/g0dm0d/uptime/internal/store"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type HeartbeatStore struct {
	writeAPI api.WriteAPIBlocking
	queryAPI api.QueryAPI
}

func NewHeartbeatStore(writeAPI api.WriteAPIBlocking, queryAPI api.QueryAPI) HeartbeatStore {
	return HeartbeatStore{
		writeAPI: writeAPI,
		queryAPI: queryAPI,
	}
}

func (s *HeartbeatStore) SaveTick(opts store.Tick) error {
	result, _ := json.Marshal(opts)
	fields := map[string]interface{}{
		"result": result,
	}

	pointname := fmt.Sprint(opts.MonitorID)

	point := write.NewPoint(pointname, map[string]string{}, fields, time.Now())
	err := s.writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		return err
	}

	return nil
}

func (s *HeartbeatStore) GetTickHistory(monitorID string, count, timestamp int) ([]store.Heartbeat, error) {
	ticks := []store.Heartbeat{}

	query := fmt.Sprintf(`from(bucket: "uptime")
					|> range(start: %d)
					|> sort(columns: ["_time"], desc: true)
          |> filter(fn: (r) => r["_measurement"] == "%s")
    			|> filter(fn: (r) => r["_field"] == "result")
    			|> limit(n:%d, offset: 0)`, timestamp, monitorID, count)

	results, err := s.queryAPI.Query(context.Background(), query)
	if err != nil {
		return ticks, err
	}

	for results.Next() {
		var tick store.Tick
		json.Unmarshal([]byte(results.Record().ValueByKey("_value").(string)), &tick)
		hearbeat := store.Heartbeat{
			Date: results.Record().ValueByKey("_time").(time.Time),
			Tick: tick,
		}
		ticks = append(ticks, hearbeat)
	}

	if err := results.Err(); err != nil {
		return ticks, err
	}

	return ticks, err
}
