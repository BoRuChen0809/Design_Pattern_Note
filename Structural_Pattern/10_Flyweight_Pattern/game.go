package main

type game struct {
	homePlayers []*player
	awayPlayers []*player
}

func newGame() *game {
	return &game{
		homePlayers: make([]*player, 1),
		awayPlayers: make([]*player, 1),
	}
}

func (g *game) addHomePlayer(jerseyType string) {
	player := newPlayer("Home", jerseyType)
	g.homePlayers = append(g.homePlayers, player)
}

func (g *game) addAwayPlayer(jerseyType string) {
	player := newPlayer("Away", jerseyType)
	g.awayPlayers = append(g.awayPlayers, player)
}
