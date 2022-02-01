//排
package main

import "fmt"

type platoon struct {
	children []army
	Name     string
	People   int
}

func (p *platoon) showInfo() {
	fmt.Printf("the platoon name is %s, there are %d people in here. ", p.getNmae(), p.getPeople())
	fmt.Printf("there are %d squards in this platoon, detail:\n", len(p.children))
	for _, a := range p.children {
		fmt.Print("。")
		a.showInfo()
	}
}

func (p *platoon) getLevel() int {
	return 2
}

func (p *platoon) setPeople() {
	num := 0
	for _, a := range p.children {
		num += a.getPeople()
	}
	p.People = num
}

func (p *platoon) getPeople() int {
	p.setPeople()
	return p.People
}

func (p *platoon) getNmae() string {
	return p.Name
}

func (p *platoon) setName(name string) {
	p.Name = name
}

func (p *platoon) add(a army) {
	if p.getLevel()-a.getLevel() != 1 {
		fmt.Println("this operation is invaild")
	} else {
		p.children = append(p.children, a)
	}
}
