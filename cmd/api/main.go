package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zahidhasanzh/olx-api/internal/config"
	"github.com/zahidhasanzh/olx-api/internal/db"
	"github.com/zahidhasanzh/olx-api/internal/handlers"
)

func main() {
	cnf := config.MustLoad()
	_, err := db.Connect(cnf.DatabaseUrl)
	if err != nil{
		log.Fatalf("main.db.connect: %v", err)
	}

	fmt.Println("database connected")
	fmt.Println("starting olx server...")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Healthz)

	srv := http.Server{
		Addr:         ":" + cnf.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	fmt.Printf("server is listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
