package postgresql

import (
	"database/sql"

	"github.com/g0dm0d/uptime/internal/store"
	"github.com/lib/pq"
)

type MonitorStore struct {
	db *sql.DB
}

func NewMonitorStore(db *sql.DB) store.MonitorStore {
	return &MonitorStore{
		db: db,
	}
}

func (s *MonitorStore) AddMonitor(opts store.CreateMonitorOpts) (id int, err error) {
	row := s.db.QueryRow("SELECT * FROM add_monitor($1, $2, $3, $4, $5, $6)",
		opts.Hostname, opts.Interval, opts.Protocol, opts.Addr, opts.Port, pq.Array(opts.Tags))

	err = row.Scan(&id)

	return id, err
}

func (s *MonitorStore) GetMonitor(id int) (store.Monitor, error) {
	var monitor store.Monitor
	var asd string
	req := s.db.QueryRow("SELECT * FROM get_monitor($1)", id)
	err := req.Scan(&asd, &monitor.Hostname, &monitor.Interval,
		&monitor.Protocol, &monitor.Addr, &monitor.Port, pq.Array(&monitor.Tags))

	return monitor, err
}

func (s *MonitorStore) GetAllMonitor() ([]store.Monitor, error) {
	var monitors []store.Monitor

	rows, err := s.db.Query("SELECT id, hostname, interval, protocol, addr, port, tags FROM monitors")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var monitor store.Monitor

		err = rows.Scan(&monitor.ID, &monitor.Hostname, &monitor.Interval,
			&monitor.Protocol, &monitor.Addr, &monitor.Port, pq.Array(&monitor.Tags))

		if err != nil {
			return nil, err
		}
		monitors = append(monitors, monitor)
	}

	return monitors, err
}
