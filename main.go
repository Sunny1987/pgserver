package main

import (
	"context"
	"github.com/common-nighthawk/go-figure"
	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pgserver/handlers"
	"time"
)

func main() {
	//heroku related updates for port
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	l := log.New(os.Stdout, "PGServ:", log.LstdFlags)
	parseHandler := handlers.GetNewLogger(l)
	serverMux := mux.NewRouter()

	//Register the handlers to the server mux
	getRouter := serverMux.Methods("GET").Subrouter()
	getRouter.HandleFunc("/users", parseHandler.GetUsers)
	getRouter.HandleFunc("/scans", parseHandler.GetScans)
	postRouter := serverMux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/addscan", parseHandler.AddScan)
	deleteRouter := serverMux.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/deletescan/{id}", parseHandler.DeleteAScan)
	//CORS
	ch := goHandlers.CORS(goHandlers.AllowedOrigins([]string{"*"}))

	//heroku related updates
	port = ":" + port

	//local test
	//port := ":8080"

	prodServer := &http.Server{
		Addr:         port,
		Handler:      ch(serverMux),
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     l,
	}

	go func() {
		myFigure := figure.NewFigure("PGServer", "", true)
		myFigure.Print()
		l.Println("version: 1.0.0")
		l.Println("Author: Sabyasachi Roy")
		l.Println("Starting PG server...")
		if err := prodServer.ListenAndServe(); err != nil {
			l.Printf("Error starting server: %v", err)
			os.Exit(1)
		}

	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan

	l.Println("Stopping server as per user interrupt", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	prodServer.Shutdown(tc)
}
