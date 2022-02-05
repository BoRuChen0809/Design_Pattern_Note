package main

import "fmt"

func main() {
	game := newGame()

	game.addHomePlayer(homeJerseyType)
	game.addHomePlayer(homeJerseyType)
	game.addHomePlayer(homeJerseyType)
	game.addHomePlayer(homeJerseyType)

	game.addAwayPlayer(awayJerseyType)
	game.addAwayPlayer(awayJerseyType)
	game.addAwayPlayer(awayJerseyType)
	game.addAwayPlayer(awayJerseyType)

	jerseyFactory := getJerseyFactory()
	for jerseyType, jersey := range jerseyFactory.jerseymap {
		fmt.Printf("JerseyColorType : %s\t JerseyColor : %s\n", jerseyType, jersey.getColor())
	}
}
