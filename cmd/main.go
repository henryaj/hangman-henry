package main

import (
	"fmt"

	ui "github.com/gizak/termui"
	"github.com/henryaj/hangman-henry/cmd/client"
	"github.com/henryaj/hangman-henry/game"
)

func main() {
	serverURI := "http://localhost:8000"

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	fmt.Println(`Welcome to hangman!

What do you want to do now?

l - list all games
n - start a new game

> `)

	for e := range ui.PollEvents() {
		switch e.ID {
		case "q", "<C-c>":
			return
		case "l":
			fmt.Println("")
			fmt.Println("Getting list of games...")
			fmt.Println("")

			fmt.Println(client.ListGames(serverURI))
			fmt.Println("")
			fmt.Println("> ")
			break
		case "n":
			fmt.Println("")
			fmt.Println("Starting a new game...")

			id, game := client.CreateNewGame(serverURI)

			playGame(serverURI, id, game)
		default:
			break
		}
	}
}

func playGame(serverURI, id string, game *game.Game) {
	fmt.Println("")
	fmt.Println("Hit Ctrl-C to exit the game")
	fmt.Println(game)
	fmt.Println("")
	fmt.Println("Press a letter to play it")
	fmt.Println(">")

	for e := range ui.PollEvents() {
		switch e.ID {
		case "<C-c>":
			fmt.Println(`What do you want to do now?

l - list all games
n - start a new game

> `)
			return
		default:
			_, game, err := client.MakeGameMove(serverURI, id, e.ID)

			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(game)
			}
		}
	}
}
