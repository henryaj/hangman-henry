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

// GamesMap is a hash map of all known games, keyed
// by game ID.
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

// APIServer is a singleton representing the API server.
type APIServer struct {
	games  GamesMap
	lock   sync.RWMutex
	logger *log.Logger
}

// GameWithID is a wrapper for a Game struct and its unique ID.
type GameWithID struct {
	ID   string
	Game *game.Game
}

// NewAPIServer initialises and returns an API Server.
func NewAPIServer() *APIServer {
	return &APIServer{
		games:  make(map[string]*game.Game),
		lock:   sync.RWMutex{},
		logger: log.New(os.Stdout, "http: ", log.LstdFlags),
	}
}

// CreateNewGame creates a game with a random word and returns it,
// along with its unique ID.
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

// ListGames lists all ongoing and past games.
func (a *APIServer) ListGames() map[string]*game.Game {
	return a.games
}

// GetGame finds a game by its ID and returns it, or an error if not found.
func (a *APIServer) GetGame(id string) (*game.Game, error) {
	a.lock.RLock()
	game, ok := a.games[id]
	a.lock.RUnlock()

	if !ok {
		return nil, fmt.Errorf("game not found")
	}
	return game, nil
}

// MakeGameMove makes a move for the specified game, returning a summary of
// that game.
func (a *APIServer) MakeGameMove(id string, letter string) (*GameWithID, error) {
	game, err := a.GetGame(id)

	if err != nil {
		return nil, err
	}

	if len(letter) != 1 {
		return nil, fmt.Errorf("invalid play: play a single letter")
	}

	err = game.Try(letter)
	if err != nil {
		return nil, err
	}

	return &GameWithID{
		ID:   id,
		Game: game,
	}, nil
}
