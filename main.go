package main

import (
	"mjgf/mjgf"
)

func main() {
	var guessGame = NewGuessGame(2)
	var game mjgf.MJGF
	game.Register(guessGame)
	game.Start()
}
