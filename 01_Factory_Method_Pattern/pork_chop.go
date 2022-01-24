package main

type pork_chop struct {
	food
}

func newPork_Chop() product {
	return &pork_chop{food: food{
		name:     "pork chop",
		describe: "It is a pork chop.",
	}}
}
