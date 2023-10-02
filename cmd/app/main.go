package main

import (
	"fmt"
	"log"

	"github.com/g0dm0d/uptime/internal/config"
	"github.com/g0dm0d/uptime/internal/server"
	"github.com/g0dm0d/uptime/internal/server/socket"
	"github.com/g0dm0d/uptime/internal/service"
	"github.com/g0dm0d/uptime/internal/store/influxdb"
	"github.com/g0dm0d/uptime/internal/store/postgresql"
	"github.com/g0dm0d/uptime/internal/uptime"
	"github.com/g0dm0d/uptime/pkg/cron"
	"github.com/g0dm0d/uptime/pkg/jwtmanager"
)

func main() {
	config := config.New()

	fmt.Printf("Version is: %s.\nListening %s:%d\n", config.App.Version, config.App.Addr, config.App.Port)

	// Init jwt service
	jwt := jwtmanager.New(config.App.Secret)

	// Init DB
	writeAPI, queryAPI, err := influxdb.New(influxdb.NewParams{
		Token:  config.InfluxDB.Token,
		Org:    config.InfluxDB.Org,
		Bucket: config.InfluxDB.Bucket,
		Addr:   config.InfluxDB.Addr,
		Port:   config.InfluxDB.Port,
	})

	if err != nil {
		log.Fatal(err)
	}

	db, err := postgresql.New(config.PostgresDB.DNS)

	if err != nil {
		log.Fatal(err)
	}

	// Init store
	hertbeatStore := influxdb.NewHeartbeatStore(*writeAPI, *queryAPI)
	userStore := postgresql.NewUserStore(db)
	monitorStore := postgresql.NewMonitorStore(db)

	// Init cron module for uptime
	c := cron.NewCron()

	// Init WebSocket
	ws := socket.New()

	// Init and run uptime ping service
	u := uptime.New(*c, &hertbeatStore, monitorStore, ws)
	err = u.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Init services
	services := service.New(service.Opts{
		UserStore:      userStore,
		MonitorStore:   monitorStore,
		HeartbeatStore: &hertbeatStore,
		JWT:            *jwt,
		Uptime:         *u,
	})

	server := server.NewServer(&server.Config{
		Addr:       config.App.Addr,
		Port:       config.App.Port,
		Service:    services,
		JwtManager: jwt,
		WebSocket:  ws,
	})

	server.SetupRouter()

	log.Print("Server up and running.")

	err = server.RunServer()
	if err != nil {
		log.Panic(err)
	}
}
