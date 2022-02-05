package main

type player struct {
	jersey     jersey
	playerType string
	number     int
	position   string
}

func newPlayer(playerType string, jerseyType string) *player {
	jersey, _ := getJerseyFactory().getJerseyByType(jerseyType)
	return &player{playerType: playerType, jersey: jersey}
}

func (p *player) setPosition(position string) {
	p.position = position
}

func (p *player) setNumber(num int) {
	p.number = num
}
