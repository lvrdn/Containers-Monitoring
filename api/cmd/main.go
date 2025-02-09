package main

import (
	"api/pkg/config"
	"api/pkg/handler"
	"api/pkg/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("get config error: [%s]\n", err.Error())
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DBusername,
		cfg.DBpassword,
		cfg.DBhost,
		cfg.DBname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("open sql connection failed, error: [%s]\n", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failed, error: [%s]\n", err.Error())
	}

	containerHandler := handler.NewContainerHandler(postgres.NewStorage(db, len(cfg.Addresses)))

	err = containerHandler.InitData(cfg.Addresses)
	if err != nil {
		log.Fatalf("add to db container addresses error: [%s]\n", err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/containers", containerHandler.UpdateData)
	mux.HandleFunc("GET /api/containers", containerHandler.ShowData)

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
