package pinger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

type Pinger struct {
	PingAddr  string
	AddrAPI   string
	MethodAPI string
	Timeout   time.Duration
	Frequency time.Duration
}

type Info struct {
	Addr  string    `json:"addr"`
	Alive bool      `json:"alive"`
	Time  time.Time `json:"last_ping_time"`
}

func (p *Pinger) Run(ctx context.Context) {

	ticker := time.NewTicker(p.Frequency)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			time := time.Now()
			alive := p.ping()

			info := &Info{
				Addr:  p.PingAddr,
				Alive: alive,
				Time:  time,
			}
			err := p.send(p.AddrAPI, p.MethodAPI, info)
			if err != nil {
				log.Printf("send ping info error: dest [%s], error text [%s]", p.AddrAPI, err.Error())
			}
		}
	}
}

func (p *Pinger) ping() bool {

	conn, err := net.DialTimeout("tcp", p.PingAddr, p.Timeout)
	if conn != nil {
		conn.Close()
	}

	var alive bool

	if err != nil {
		log.Printf("ping failed: addr [%s], error text [%s]\n", p.PingAddr, err.Error())
		alive = false
	} else {
		alive = true
	}

	return alive
}

func (p *Pinger) send(addr, method string, info *Info) error {

	dataToSend, err := json.Marshal(info)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, addr, bytes.NewBuffer(dataToSend))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: expected [%d], got [%d]", http.StatusOK, resp.StatusCode)
	}

	return nil
}
