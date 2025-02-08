package container

import "time"

type Container struct {
	Addr            string     `json:"addr"`
	Alive           *bool      `json:"alive,omitempty"`
	LastPing        *time.Time `json:"last_ping_time"`
	LastSuccessPing *time.Time `json:"last_alive_time"`
}
