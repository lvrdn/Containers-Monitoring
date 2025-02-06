package pinger

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Pinger struct {
	PingAddr  string
	AddrAPI   string
	MethodAPI string
	Timeout   time.Duration
}

func (p *Pinger) Ping(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	time := time.Now()
	conn, err := net.DialTimeout("tcp", p.PingAddr, p.Timeout)
	if conn != nil {
		conn.Close()
	}

	var alive bool

	if err != nil {
		//log.Printf("ping failed: addr [%s]\n", p.PingAddr)
		alive = false
	} else {
		alive = true
	}

	err = SendInfo(p.AddrAPI, p.MethodAPI, p.PingAddr, alive, time)
	if err != nil {
		log.Printf("send ping info error: dest [%s], error text [%s]", p.AddrAPI, err.Error())
	}

}

func SendInfo(addr, method, pingedAddr string, alive bool, time time.Time) error {

	dataToSend := []byte(fmt.Sprintf(`{"addr":"%s","alive":%v,"last_ping_time":"%s"}`, pingedAddr, alive, time.Format("2006-01-02T15:04:05Z07:00")))

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
