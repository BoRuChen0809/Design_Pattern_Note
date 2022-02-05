package main

type awayJersey struct {
	color string
}

func (a *awayJersey) getColor() string {
	return a.color
}

func newAwayJersey() *awayJersey {
	return &awayJersey{color: "dark"}
}
