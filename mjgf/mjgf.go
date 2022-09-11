package mjgf

import "fmt"

type Status struct {
	IsTurn bool
	Status map[string]string
}

type Move struct {
	Player *Player
	Move   map[string]string
}

type Game interface {
	HasEnded() bool
	GetMaxPlayers() int
	GetPlayers() *[]*Player

	StartGame(*[]*Player)
	NewMove(Move)
}

// Player definition
type Player struct {
	Id     string
	IsTurn bool
	Status Status
}

func (p *Player) reportStatus() {
	fmt.Printf("Status report for Player %v:\n", p.Id)
	fmt.Printf("\tisTurn: %t\n", p.IsTurn)
	for k, v := range p.Status.Status {
		fmt.Printf("\t%v: %v\n", k, v)
	}
}

// Framework definition
type MJGF struct {
	Game Game
}

func (mjgf *MJGF) Register(g Game) {
	mjgf.Game = g
	fmt.Println("Game registered!")
}

func (mjgf *MJGF) Start() {
	g := mjgf.Game
	// Lobby logic
	fmt.Println("Starting Lobby:")
	var players []*Player
	maxPlayers := g.GetMaxPlayers()
	for player := 0; player < maxPlayers; player++ {
		var id string
		fmt.Printf("Enter ID for Player %v: ", player)
		fmt.Scan(&id)
		players = append(players, &Player{Id: id, IsTurn: false, Status: Status{}})
	}
	g.StartGame(&players)
	for _, player := range *g.GetPlayers() {
		player.reportStatus()
	}

	// Game loop logic
	for {
		if mjgf.Game.HasEnded() {
			break
		}
		for _, player := range *g.GetPlayers() {
			if player.IsTurn {
				// wait for move from players
				//! Separation of concerns (Guess logic no va aqui)
				var guess string
				fmt.Printf("Input for Player %v-> ", player.Id)
				fmt.Scan(&guess)
				move := Move{
					player,
					map[string]string{
						"guess": guess,
					},
				}

				// update game state
				g.NewMove(move)
				if mjgf.Game.HasEnded() {
					break
				}

				// send status to players
				for _, player := range *g.GetPlayers() {
					player.reportStatus()
				}
			}
		}
	}

	// End game logic
	fmt.Println("Game has finished!")
}
