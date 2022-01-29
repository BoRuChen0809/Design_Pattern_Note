package main

import "fmt"

type Chef struct {
}

func newChef() *Chef {
	return &Chef{}
}

func (c *Chef) cook() {
	fmt.Println("下廚烹飪......")
}
