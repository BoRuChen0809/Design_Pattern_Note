//連
package main

import "fmt"

type company struct {
	children []army
	Name     string
	People   int
}

func (c *company) showInfo() {
	fmt.Printf("the company name is %s, there are %d people in here. ", c.getNmae(), c.getPeople())
	fmt.Printf("there are %d platoons in this company, detail:\n", len(c.children))
	for _, a := range c.children {
		fmt.Print("。")
		a.showInfo()
	}
}

func (c *company) getLevel() int {
	return 3
}

func (c *company) setPeople() {
	num := 0
	for _, a := range c.children {
		num += a.getPeople()
	}
	c.People = num
}

func (c *company) getPeople() int {
	c.setPeople()
	return c.People
}

func (c *company) getNmae() string {
	return c.Name
}

func (c *company) setName(name string) {
	c.Name = name
}

func (c *company) add(a army) {
	if c.getLevel()-a.getLevel() != 1 {
		fmt.Println("this operation is invaild")
	} else {
		c.children = append(c.children, a)
	}
}
