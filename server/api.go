package server

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"

	"github.com/henryaj/hangman-henry/game"
	uuid "github.com/satori/go.uuid"
)

type GamesMap map[string]*game.Game

func (g GamesMap) String() string {
	var gamesList []string

	for id, game := range g {
		var lettersAttempted = "none"

		if len(game.LettersAttempted) > 0 {
			lettersAttempted = strings.Join(game.LettersAttempted, "")
		}

		gameString := fmt.Sprintf(
			"%s | %d | %s",
			id,
			game.AttemptsRemaining,
			lettersAttempted,
		)

		gamesList = append(gamesList, gameString)
	}

	return "ALL GAMES\n" +
		"---------\n" +
		strings.Join(gamesList, "\n")
}

type APIServer struct {
	games  GamesMap
	lock   sync.RWMutex
	logger *log.Logger
}

type GameSummary struct {
	ID   string
	Game *game.Game
}

func NewAPIServer() *APIServer {
	return &APIServer{
		games:  make(map[string]*game.Game),
		lock:   sync.RWMutex{},
		logger: log.New(os.Stdout, "http: ", log.LstdFlags),
	}
}

func (a *APIServer) CreateNewGame() (string, *game.Game) {
	words := []string{"alabaster", "lobster", "loofah"}
	n := rand.Intn(len(words))
	word := words[n]

	g := game.NewGame(word, 10)
	id := uuid.Must(uuid.NewV4(), nil).String()

	a.lock.Lock()
	a.games[id] = g
	a.lock.Unlock()

	return id, g
}

func (a *APIServer) ListGames() map[string]*game.Game {
	return a.games
}

func (a *APIServer) GetGame(id string) (*game.Game, error) {
	a.lock.RLock()
	game, ok := a.games[id]
	a.lock.RUnlock()

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
