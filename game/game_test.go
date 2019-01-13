package game_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/henryaj/hangman-henry/game"
)

var _ = Describe("Game", func() {
	var game *Game

	BeforeEach(func() {
		game = NewGame("abc", 1)
	})

	Describe("Try", func() {
		When("there are still attempts remaining", func() {
			It("records the tried letter", func() {
				game.Try("a")
				Expect(game.LettersAttempted).To(ContainElement("a"))
			})

			It("returns an error if the letter has already been tried", func() {
				game.Try("a")
				err := game.Try("a")
				Expect(err).To(MatchError(fmt.Errorf("letter has already been played")))
			})

			When("the letter is not contained in the target word", func() {
				It("reduces the number of remaining attempts by one", func() {
					game.Try("z")
					Expect(game.AttemptsRemaining).To(Equal(0))
				})
			})

			When("the letter is contained in the target word", func() {
				It("does not reduce the number of remaining attempts", func() {
					game.Try("a")
					Expect(game.AttemptsRemaining).To(Equal(1))
				})
			})
		})

		When("there are no attempts remaining", func() {
			It("returns an error", func() {
				err := game.Try("q")
				Expect(err).To(MatchError(fmt.Errorf("game over - the word was 'abc'")))

				Expect(game.Try("y")).To(MatchError(fmt.Errorf("game's over, stop playing")))
			})
		})
	})

	Describe("InProgress", func() {
		It("returns true when the game is in progress, false otherwise", func() {
			Expect(game.InProgress()).To(BeTrue())

			game.Try("q")
			Expect(game.InProgress()).To(BeFalse())
		})
	})

	Describe("Won", func() {
		It("returns true when the game has been won, false otherwise", func() {
			Expect(game.Won()).To(BeFalse())

			game.Try("a")
			game.Try("b")
			game.Try("c")

			Expect(game.Won()).To(BeTrue())
		})
	})

	Describe("Lost", func() {
		It("returns true when the game has been lost, false otherwise", func() {
			Expect(game.Lost()).To(BeFalse())

			game.Try("z")
			Expect(game.Lost()).To(BeTrue())
		})
	})
})
