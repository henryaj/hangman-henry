package server_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/henryaj/hangman-henry/server"
)

var _ = Describe("API", func() {
	var apiServer *APIServer

	BeforeEach(func() {
		apiServer = NewAPIServer()
	})

	Describe("CreateNewGame", func() {
		It("creates a new game and returns it", func() {
			_, game := apiServer.CreateNewGame()

			Expect(string(game.GameState)).To(Equal("in progress"))
		})
	})

	Describe("ListGames", func() {
		It("returns a summary of the games in progress", func() {
			id1, _ := apiServer.CreateNewGame()
			id2, _ := apiServer.CreateNewGame()
			id3, _ := apiServer.CreateNewGame()

			games := apiServer.ListGames()
			Expect(games).To(HaveLen(3))

			var ids []string

			for id, _ := range games {
				ids = append(ids, id)
			}

			Expect(ids).To(ConsistOf(id1, id2, id3))
		})
	})

	Describe("GetGame", func() {
		It("returns the current game state", func() {
			id, _ := apiServer.CreateNewGame()

			foundGame, err := apiServer.GetGame(id)
			Expect(err).NotTo(HaveOccurred())
			Expect(foundGame).NotTo(BeNil())
		})

		It("returns an error if the game does not exist", func() {
			_, err := apiServer.GetGame("dave")
			Expect(err).To(MatchError(fmt.Errorf("game not found")))
		})
	})

	Describe("MakeGameMove", func() {
		It("returns the status of the game", func() {
			id, _ := apiServer.CreateNewGame()

			status, err := apiServer.MakeGameMove(id, "z")
			Expect(err).NotTo(HaveOccurred())

			Expect(status.Game.LettersAttempted).To(ConsistOf("z"))
		})
	})
})
