package main

import "fmt"

type iFactory interface {
	makeShirt() iShirt
	makeShoe() iShoe
}

func getSportsFactory(brand string) (iFactory, error) {
	if brand == "UA" {
		return &UA{}, nil
	}

	if brand == "Nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
