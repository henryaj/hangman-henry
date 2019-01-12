package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/henryaj/hangman-henry/server"
)

func main() {
	api := server.NewAPIServer()
	r := mux.NewRouter()

	handleRoute("GET", "/games", api.ListGameHandler, r)
	handleRoute("POST", "/games", api.NewGameHandler, r)
	handleRoute("GET", "/games/{id}", api.GetGameHandler, r)
	handleRoute("POST", "/games/{id}/{letter}", api.MakeGameMoveHandler, r)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		ErrorLog:     logger,
	}

	log.Fatal(srv.ListenAndServe())
}

func handleRoute(method string, route string, handleFunc http.HandlerFunc, r *mux.Router) {
	r.Handle(route,
		handlers.LoggingHandler(os.Stdout, http.HandlerFunc(handleFunc)),
	).Methods(method)
}
