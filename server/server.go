package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/henryaj/hangman-henry/game"
)

type GameWithID struct {
	ID   string    `json:"id"`
	Game game.Game `json:"game"`
}

func (a *APIServer) NewGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, game := a.CreateNewGame()

	resp, err := json.Marshal(GameWithID{ID: id, Game: *game})
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)

	io.WriteString(w, string(resp))
}

func (a *APIServer) ListGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(a.ListGames())
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(resp))
}

func (a *APIServer) GetGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	game, err := a.GetGame(id)
	if err != nil {
		handleError(err, w)
		return
	}

	resp, err := json.Marshal(game)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(resp))
}

func (a *APIServer) MakeGameMoveHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	letter := vars["letter"]

	gameSummary, err := a.MakeGameMove(id, letter)
	if err != nil {
		handleError(err, w)
		return
	}

	resp, err := json.Marshal(gameSummary)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(resp))
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, fmt.Sprintf(
		`{"error": "%s"}`, err,
	))
}
