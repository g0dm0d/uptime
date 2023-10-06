package influxdb

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type NewParams struct {
	Token  string
	Org    string
	Bucket string
	Addr   string
	Port   int
}

func New(param NewParams) (*api.WriteAPIBlocking, *api.QueryAPI, error) {
	client := influxdb2.NewClient(fmt.Sprint(param.Addr, ":", param.Port), param.Token)
	ok, err := client.Ping(context.Background())
	if err != nil || !ok {
		return nil, nil, err
	}

	wrtieAPI := client.WriteAPIBlocking(param.Org, param.Bucket)
	queryAPI := client.QueryAPI(param.Org)
	return &wrtieAPI, &queryAPI, nil
}
