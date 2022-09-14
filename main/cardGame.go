package main

import (
	"cardGame"
	"fmt"
)

func main() {

	// initialize Deck
	var newGame cardGame.TeenPati
	var playerName string
	playerNumber := 0

	fmt.Print("Enter Player Name (without space) : ")
	fmt.Scan(&playerName)

	for (playerNumber < 1) || (playerNumber > 3) {
		fmt.Print("Enter Number of AI (min 1 , max 3) : ")
		fmt.Scan(&playerNumber)
	}

	newGame.InitGame(playerName, playerNumber)

	newGame.StartGame()

}
