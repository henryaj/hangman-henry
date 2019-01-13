# hangman

A simple hangman game written in Go. The client and the server communicate using JSON over HTTP; the API is stateless by design.

## Running

Build the server binary with `go build -o server`. Build the client with `cd cmd; go build -o client`. Then run the binaries.

Or, just run `go run .` and `cd cmd; go run .` in two terminal windows.

## Tests

Run tests with `ginkgo -r` (you'll need to `go get -u github.com/onsi/ginkgo` first).

## Caveats

Users can view a list of in-progress/completed games, but currently can't play them; they can only create new games.

The client isn't perfect – the logic for handling keypresses is maybe a bit eccentric, and it could do with tests (though it's tricky to test a termbox). For now, the test suite just ensures that binary can be build without error.
