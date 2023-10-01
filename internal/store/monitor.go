package store

import "database/sql"

type Protocol string

const (
	HTTP Protocol = "http"
	TCP  Protocol = "tcp"
)

func (p Protocol) IsValid() bool {
	switch p {
	case HTTP, TCP:
		return true
	}
	return false
}

type Monitor struct {
	ID       int
	Hostname string
	Interval int
	Protocol Protocol
	Addr     string
	Port     sql.NullInt16
	Tags     []string
}

type CreateMonitorOpts struct {
	Hostname string
	Interval int
	Protocol string
	Addr     string
	Port     interface{}
	Tags     []string
}

type MonitorStore interface {
	AddMonitor(opts CreateMonitorOpts) error
	GetMonitor(id int) (Monitor, error)
	GetAllMonitor() ([]Monitor, error)
}
