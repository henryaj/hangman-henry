package server

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/henryaj/hangman-henry/game"
	uuid "github.com/satori/go.uuid"
)

type GamesMap map[string]*game.Game

type APIServer struct {
	games  GamesMap
	logger *log.Logger
}

type GameSummary struct {
	ID   string
	Game *game.Game
}

func NewAPIServer() *APIServer {
	return &APIServer{
		games:  make(map[string]*game.Game),
		logger: log.New(os.Stdout, "http: ", log.LstdFlags),
	}
}

func (a *APIServer) CreateNewGame() string {
	words := []string{"alabaster", "lobster", "loofah"}
	n := rand.Intn(len(words))
	word := words[n]

	g := game.NewGame(word, 10)
	id := uuid.Must(uuid.NewV4(), nil).String()

	a.games[id] = g

	return id
}

func (a *APIServer) ListGames() map[string]*game.Game {
	return a.games
}

func (a *APIServer) GetGame(id string) (*game.Game, error) {
	game, ok := a.games[id]

	if !ok {
		return nil, fmt.Errorf("game not found")
	}
	return game, nil
}

func (a *APIServer) MakeGameMove(id string, letter string) (*GameSummary, error) {
	game, err := a.GetGame(id)

	if err != nil {
		return nil, err
	}

	//TODO: prevent user from playing more than one letter!
	err = game.Try(letter)
	if err != nil {
		return nil, err
	}

	return &GameSummary{
		ID:   id,
		Game: game,
	}, nil
}
