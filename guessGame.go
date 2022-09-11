package main

import (
	"fmt"
	"mjgf/mjgf"
	"strconv"
)

type GuessGame struct {
	maxPlayers int
	state      map[string]string
	players    *[]*mjgf.Player
	hasEnded   bool
}

func NewGuessGame(maxPlayers int) *GuessGame {
	return &GuessGame{
		maxPlayers: maxPlayers,
		state:      map[string]string{},
		players:    &[]*mjgf.Player{},
	}
}

func (g *GuessGame) StartGame(players *[]*mjgf.Player) {
	// Initialize players
	g.players = players

	// Initialize game state
	g.state = make(map[string]string)
	g.state["magic_character"] = "j"
	g.state["guessed"] = "false"
	g.state["current_player_index"] = "0"
	g.state["current_player_id"] = (*players)[0].Id
	fmt.Printf("Started game with %v players\n", len(*g.players))

	// Report initial status to players
	for _, player := range *g.players {
		player.IsTurn = g.state["current_player_id"] == player.Id
		player.Status = mjgf.Status{
			Status: map[string]string{"guessed": g.state["guessed"]},
		}
	}
}

func (g *GuessGame) NewMove(move mjgf.Move) {
	if g.state["current_player_id"] == move.Player.Id {
		// Update State
		g.state["guessed"] = strconv.FormatBool(move.Move["guess"] == g.state["magic_character"])
		currentPlayerIndex, _ := strconv.Atoi(g.state["current_player_index"])
		nextCurrentPlayerIndex := (currentPlayerIndex + 1) % g.maxPlayers
		g.state["current_player_index"] = strconv.Itoa(nextCurrentPlayerIndex)
		g.state["current_player_id"] = (*g.players)[nextCurrentPlayerIndex].Id
		g.hasEnded = g.state["guessed"] == "true"

		// Update players status
		for _, player := range *g.players {
			player.IsTurn = g.state["current_player_id"] == player.Id
			player.Status = mjgf.Status{
				Status: map[string]string{"guessed": g.state["guessed"]},
			}
		}
	}
}

func (g *GuessGame) GetMaxPlayers() int {
	return g.maxPlayers
}

func (g *GuessGame) HasEnded() bool {
	return g.hasEnded
}

func (g *GuessGame) GetPlayers() *[]*mjgf.Player {
	return g.players
}
