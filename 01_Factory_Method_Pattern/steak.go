package main

type steak struct {
	food
}

func newSteak() product {
	return &steak{food: food{
		name:     "SteaKKKKKKKKKKK",
		describe: "Medium",
	}}
}
