package main

import (
	"log"
	"os"
	"os/signal"
	"pinger/pkg/config"
	"pinger/pkg/pinger"
	"sync"
	"syscall"
	"time"
)

type Str struct {
	Name string
}

func main() {

	cfg, err := config.GetConfig("./app.env")
	if err != nil {
		log.Fatalf("get config failed: [%s]\n", err.Error())
	}

	ticker := time.NewTicker(cfg.PingFrequency)
	wg := &sync.WaitGroup{}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	workers := make([]*pinger.Pinger, len(cfg.Addresses))

	for i := range cfg.Addresses {
		workers[i] = &pinger.Pinger{
			PingAddr:  cfg.Addresses[i],
			AddrAPI:   cfg.AddrAPI,
			MethodAPI: cfg.MethodAPI,
			Timeout:   cfg.PingTimeout,
		}
	}

LOOP:
	for {
		select {
		case <-ticker.C:
			for _, worker := range workers {
				go worker.Ping(wg)
			}

		case <-stop:
			wg.Wait()
			break LOOP

		default:
			continue
		}
	}

}
