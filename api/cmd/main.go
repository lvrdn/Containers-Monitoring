package main

import (
	"api/pkg/config"
	"api/pkg/container"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("get config error: [%s]\n", err.Error())
	}

	// dsn := fmt.Sprintf(
	// 	"postgres://%s:%s@%s/%s?sslmode=disable",
	// 	cfg.DBusername,
	// 	cfg.DBpassword,
	// 	cfg.DBhost,
	// 	cfg.DBname,
	// )

	// db, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	log.Fatalf("open sql connection failed, error: [%s]\n", err.Error())
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("db  ping failed, error: [%s]\n", err.Error())
	// }

	containerHandler := container.NewContainerHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/containers", containerHandler.ReceiveData)

	server := http.Server{
		Addr:    ":" + cfg.HTTPport,
		Handler: mux,
	}

	go func() {
		log.Println("start server on:", cfg.HTTPport)
		server.ListenAndServe()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	server.Shutdown(context.Background())
	log.Println("server stopped")
}
