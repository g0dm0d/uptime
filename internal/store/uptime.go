package store

import (
	"reflect"
)

type Protocol string

const (
	HTTP  Protocol = "http"
	HTTPS Protocol = "https"
	TCP   Protocol = "tcp"
)

func (p Protocol) IsValid() bool {
	switch p {
	case HTTP, TCP:
		return true
	}
	return false
}

type MonitorConfig struct {
	ID       string
	Name     string     `toml:"name"`
	Protocol Protocol   `toml:"protocol"`
	Port     int        `toml:"port"`
	Interval int        `toml:"interval"`
	Method   string     `toml:"method"`
	Addr     string     `toml:"addr"`
	Headers  [][]string `toml:"headers"`
	Body     string     `toml:"body"`
	Status   int        `toml:"valid_status"`
}

type UptimeConfig struct {
	Servers map[string]MonitorConfig `toml:"tcp"`
	HTTP    map[string]MonitorConfig `toml:"http"`
}

func (c UptimeConfig) GetList() []MonitorConfig {
	servers := []MonitorConfig{}
	protocols := reflect.ValueOf(c)
	for i := 0; i < protocols.NumField(); i++ {
		iter := protocols.Field(i).MapRange()
		for iter.Next() {
			k := iter.Key().String()
			v := iter.Value().Interface().(MonitorConfig)
			v.ID = k
			servers = append(servers, v)
		}
	}
	return servers
}
