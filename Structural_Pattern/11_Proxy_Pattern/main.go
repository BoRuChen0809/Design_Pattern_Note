package main

func main() {
	m := &me{}

	estateAgent := &EstateAgent{m}

	estateAgent.findHouse()
	estateAgent.bargain()
	estateAgent.make_deal()
}
