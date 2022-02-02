package main

import "fmt"

func main() {
	V := &veggeMania{}
	VwithCheese := &cheeseTopping{pizza: V}
	fmt.Printf("素食披薩加起司:%d元\n", VwithCheese.getPrice())

	P := &peppyPaneer{}
	PwithTomato := &tomatoTopping{pizza: P}
	PwithTomatoCheese := &cheeseTopping{pizza: PwithTomato}
	fmt.Printf("芝士爆裂披薩加番茄和起司:%d元\n", PwithTomatoCheese.getPrice())
}
