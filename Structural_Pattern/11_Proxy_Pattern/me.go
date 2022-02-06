package main

import "fmt"

type me struct{}

func (m *me) findHouse() {
	fmt.Println("找好房子ing")
}

func (m *me) bargain() {
	fmt.Println("議價......")
}

func (m *me) make_deal() {
	fmt.Println("成交!!!")
}
