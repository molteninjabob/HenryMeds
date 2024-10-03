package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	cfg "github.com/molteninjabob/HenryMeds/config"
	"github.com/molteninjabob/HenryMeds/internal/access"
	"github.com/molteninjabob/HenryMeds/internal/handlers"
	"github.com/molteninjabob/HenryMeds/internal/middleware"
)

func main() {
	ctx := context.Background()

	db, err := access.NewDbConn(ctx)
	if err != nil {
		log.Fatalf("error establishing database connection: %v", err)
	}

	mx := mux.NewRouter()
	mx.HandleFunc("/client", handlers.GetClient(ctx, db)).Methods("GET")
	mx.HandleFunc("/client", handlers.AddClient(ctx, db)).Methods("POST")
	mx.HandleFunc("/provider", handlers.GetProvider(ctx, db)).Methods("GET")
	mx.HandleFunc("/provider", handlers.AddProvider(ctx, db)).Methods("POST")
	mx.HandleFunc("/schedule", handlers.GetSchedule(ctx, db)).Methods("GET")
	mx.HandleFunc("/schedule", handlers.SetSchedule(ctx, db)).Methods("POST")
	mx.HandleFunc("/appointment/request", handlers.MakeReservation(ctx, db)).Methods("POST")
	mx.HandleFunc("/appointment/confirm", handlers.ConfirmReservation(ctx, db)).Methods("POST")

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      middleware.Auth(mx),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}

	log.Printf("starting server on port: %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("error starting server: %v", err)
	}
}
