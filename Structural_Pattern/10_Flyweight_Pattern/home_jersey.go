package main

type homeJersey struct {
	color string
}

func (h *homeJersey) getColor() string {
	return h.color
}

func newHomeJersey() *homeJersey {
	return &homeJersey{color: "light"}
}
