package main

import "fmt"

type EstateAgent struct {
	customer *me
}

func (e *EstateAgent) findHouse() {
	e.customer.findHouse()
	fmt.Println("仲介幫忙找")
}

func (e *EstateAgent) bargain() {
	e.customer.bargain()
	fmt.Println("仲介幫忙議價")
}

func (e *EstateAgent) make_deal() {
	e.customer.make_deal()
	fmt.Println("仲介幫忙處理相關手續")
}
