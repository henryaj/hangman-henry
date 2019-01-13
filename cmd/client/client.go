package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/henryaj/hangman-henry/game"
	"github.com/henryaj/hangman-henry/server"
)

// CreateNewGame sends a request to the server to start a new game,
// and returns that game and its ID.
func CreateNewGame(serverURI string) (string, *game.Game) {
	resp, err := http.Post(serverURI+"/games", "", nil)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var gameWithID server.GameWithID
	err = json.Unmarshal(body, &gameWithID)
	if err != nil {
		panic(err)
	}

	return gameWithID.ID, gameWithID.Game
}

// ListGames sends a request to the server for all known games.
func ListGames(serverURI string) server.GamesMap {
	resp, err := http.Get(serverURI + "/games")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var games server.GamesMap
	err = json.Unmarshal(body, &games)
	if err != nil {
		panic(err)
	}

	return games
}

// MakeGameMove sends a move to the server for a specified game,
// returning the resulting modified game and its ID.
func MakeGameMove(serverURI, id, move string) (string, *game.Game, error) {
	resp, err := http.Post(serverURI+"/games/"+id+"/"+move, "", nil)
	if err != nil {
		return "", &game.Game{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", &game.Game{}, err
	}

	err = checkServerError(body)
	if err != nil {
		return "", &game.Game{}, err
	}

	var gameWithID server.GameWithID
	err = json.Unmarshal(body, &gameWithID)
	if err != nil {
		return "", &game.Game{}, err
	}

	return gameWithID.ID, gameWithID.Game, nil
}

type serverError struct {
	Error string `json:"error"`
}

func checkServerError(body []byte) error {
	var srvErr serverError

	err := json.Unmarshal(body, &srvErr)
	if err != nil {
		return nil
	}

	if srvErr.Error != "" {
		return fmt.Errorf(srvErr.Error)

	}
	return nil

}
