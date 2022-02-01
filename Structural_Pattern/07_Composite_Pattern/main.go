package main

func main() {
	var a army

	squard_1 := &squard{"1", 15}
	squard_2 := &squard{"2", 20}
	squard_3 := &squard{"3", 11}

	platoon_1 := &platoon{Name: "P1"}
	platoon_1.add(squard_1)
	platoon_1.add(squard_2)
	platoon_1.add(squard_3)

	platoon_2 := &platoon{Name: "P2"}
	platoon_2.add(&squard{"4", 18})
	platoon_2.add(&squard{"7", 19})
	platoon_2.add(&squard{"5", 10})

	platoon_3 := &platoon{Name: "P3"}
	platoon_3.add(&squard{"6", 14})
	platoon_3.add(&squard{"9", 13})
	platoon_3.add(&squard{"8", 17})

	C := &company{Name: "Company"}
	C.add(platoon_2)
	C.add(platoon_3)
	C.add(platoon_1)

	a = C
	a.showInfo()
}
