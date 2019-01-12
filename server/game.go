package server

import (
	"fmt"
	"strings"
)

// Game represents a single hangman game.
type Game struct {
	Word              string
	AttemptsRemaining int
	LettersAttempted  []string
}

// NewGame returns a new game.
func NewGame(word string, attemptsRemaning int) *Game {
	return &Game{
		Word: word, AttemptsRemaining: attemptsRemaning,
	}
}

// Try plays a letter in the current game, returning an error if that
// letter has already been played.
func (g *Game) Try(letter string) error {
	if g.AttemptsRemaining == 0 {
		return fmt.Errorf("game over")
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

	return nil
}
