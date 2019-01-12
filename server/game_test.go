package server_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/henryaj/hangman-henry/server"
)

var _ = Describe("Game", func() {
	var game *Game

	BeforeEach(func() {
		game = NewGame("alabaster", 1)
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
				game.Try("z")
				game.Try("q")

				Expect(game.Try("y")).To(MatchError(fmt.Errorf("game over")))
			})
		})
	})
})
