package game

import (
	"fmt"
	"strings"
)

// Game represents a single hangman game.
type Game struct {
	Word              string    `json:"word"`
	AttemptsRemaining int       `json:"attempts_remaining"`
	LettersAttempted  []string  `json:"letters_attempted"`
	GameState         GameState `json:"game_state"`
}

type GameState int

const (
	inProgress GameState = iota + 1
	won
	lost
)

// NewGame returns a new game.
func NewGame(word string, attemptsRemaning int) *Game {
	return &Game{
		Word:              word,
		AttemptsRemaining: attemptsRemaning,
		GameState:         inProgress,
	}
}

// Try plays a letter in the current game, returning an error if that
// letter has already been played.
func (g *Game) Try(letter string) error {
	if g.GameState == lost {
		return fmt.Errorf("game already lost")
	}

	for _, char := range g.LettersAttempted {
		if char == letter {
			return fmt.Errorf("letter has already been played")
		}
	}

	g.LettersAttempted = append(g.LettersAttempted, letter)

	if !strings.Contains(g.Word, letter) {
		g.AttemptsRemaining--
	}

	if g.AttemptsRemaining == 0 {
		g.GameState = lost
		return fmt.Errorf("game over")
	}

	if sliceFullyContainsString(g.LettersAttempted, g.Word) {
		g.GameState = won
	}

	return nil
}

// InProgress returns true if the game is in progress (has not been won or lost).
func (g *Game) InProgress() bool {
	return g.GameState == inProgress
}

// Won returns true if the game has been won.
func (g *Game) Won() bool {
	return g.GameState == won
}

// Lost returns true if the game has been lost.
func (g *Game) Lost() bool {
	return g.GameState == lost
}

func sliceFullyContainsString(tried []string, target string) bool {
	sliceAsString := strings.Join(tried, "")

	for _, char := range target {
		if !strings.Contains(sliceAsString, string(char)) {
			return false
		}
	}

	return true
}
