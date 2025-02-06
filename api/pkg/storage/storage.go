package storage

import (
	"api/pkg/container"
	"database/sql"
	"time"
)

type Storage struct {
	DB      *sql.DB
	Records int
}

func NewStorage(db *sql.DB, count int) *Storage {
	return &Storage{
		DB:      db,
		Records: count,
	}
}

func (st *Storage) AddNewContainerRecord(address string) error {

	_, err := st.DB.Exec("INSERT INTO containers(address) VALUES($1) ON CONFLICT (address) DO NOTHING", address)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (st *Storage) UpdateContainerRecord(address string, time time.Time, alive bool) error {

	var err error

	if alive {
		_, err = st.DB.Exec("UPDATE containers SET last_ping=$1, last_success_ping=$2 WHERE address=$3", time, time, address)
	} else {
		_, err = st.DB.Exec("UPDATE containers SET last_ping=$1 WHERE address=$2", time, address)
	}

	if err != nil {
		return err
	}

	return nil
}

func (st *Storage) ShowContainerRecords() ([]*container.Container, error) {

	rows, err := st.DB.Query("SELECT address,last_ping,last_success_ping FROM containers ORDER BY address")
	if err != nil {
		return nil, err
	}

	containers := make([]*container.Container, 0, st.Records)

	for rows.Next() {
		var address string
		var lastPingSQL, lastSuccessPingSQL sql.NullTime
		var lastPing, lastSuccessPing *time.Time

		err := rows.Scan(&address, &lastPingSQL, &lastSuccessPingSQL)
		if err != nil {
			return nil, err
		}

		if lastPingSQL.Valid {
			lastPing = new(time.Time)
			*lastPing = lastPingSQL.Time
		}

		if lastSuccessPingSQL.Valid {
			lastSuccessPing = new(time.Time)
			*lastSuccessPing = lastSuccessPingSQL.Time
		}

		containers = append(containers, &container.Container{
			Addr:            address,
			LastPing:        lastPing,
			LastSuccessPing: lastSuccessPing,
		})

	}

	return containers, nil
}
