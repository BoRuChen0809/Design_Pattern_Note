package main

import "fmt"

type client struct {
}

func (c *client) watch(m monitor) {
	fmt.Println("Client start up monitor.")
	m.display()
}
