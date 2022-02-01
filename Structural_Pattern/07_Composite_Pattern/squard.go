//班
package main

import "fmt"

type squard struct {
	Name   string
	People int
}

func (s *squard) showInfo() {
	fmt.Printf("the squard name is %s, there are %d people in here.\n", s.getNmae(), s.getPeople())
}

func (s *squard) getLevel() int {
	return 1
}

func (s *squard) getPeople() int {
	return s.People
}

func (s *squard) setPeople(num int) {
	s.People = num
}

func (s *squard) getNmae() string {
	return s.Name
}

func (s *squard) setName(name string) {
	s.Name = name
}
