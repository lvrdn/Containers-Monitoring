package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"pinger/pkg/config"
	"pinger/pkg/pinger"
	"sync"
	"syscall"
)

type Str struct {
	Name string
}

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("get config failed: [%s]\n", err.Error())
	}

	wg := &sync.WaitGroup{}
	ctx, finish := context.WithCancel(context.Background())

	log.Println("Pinger started")

	for i := range cfg.Addresses {
		pinger := &pinger.Pinger{
			PingAddr:  cfg.Addresses[i],
			AddrAPI:   cfg.AddrAPI,
			MethodAPI: cfg.MethodAPI,
			Timeout:   cfg.PingTimeout,
			Frequency: cfg.PingFrequency,
		}

		wg.Add(1)
		go start(ctx, pinger, wg)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	finish()
	wg.Wait()

	log.Println("Pinger stopped")

}

func start(ctx context.Context, pinger *pinger.Pinger, wg *sync.WaitGroup) {
	pinger.Run(ctx)
	wg.Done()
}
