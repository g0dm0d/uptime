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

// func main() {
// 	token := os.Getenv("INFLUXDB_TOKEN")
// 	url := "http://localhost:8086"
// 	client := influxdb2.NewClient(url, token)
//
// 	org := "uptime"
// 	bucket := "<BUCKET>"
// 	writeAPI := client.WriteAPIBlocking(org, bucket)
// 	for value := 0; value < 5; value++ {
// 		tags := map[string]string{
// 			"tagname1": "tagvalue1",
// 		}
// 		fields := map[string]interface{}{
// 			"field1": value,
// 		}
// 		point := write.NewPoint("measurement1", tags, fields, time.Now())
// 		time.Sleep(1 * time.Second) // separate points by 1 second
//
// 		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
//
// func main() {
// 	p := write.NewPoint("server1",
// 		map[string]interface{}{"status": 1, "response": 200, "status": 2}, // Field values
// 		nil, // Tags (optional)
// 		nil, // Timestamp (optional, will use current time if nil)
// 	)
//
// 	// Write the data point
// 	writeAPI.WritePoint(p)
//
// 	// Close the write API
// 	writeAPI.Flush()
//
// 	// Check for errors
// 	if writeAPI.Errors() != nil {
// 		fmt.Println("Error writing data point:", writeAPI.Errors())
// 	}
// }
