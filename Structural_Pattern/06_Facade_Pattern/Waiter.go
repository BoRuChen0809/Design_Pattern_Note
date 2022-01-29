package main

import "fmt"

type Waiter struct {
}

func newWaiter() *Waiter {
	return &Waiter{}
}

func (w *Waiter) order() {
	fmt.Println("接待、入座、點菜......")
}

func (w *Waiter) service() {
	fmt.Println("上菜......")
}
