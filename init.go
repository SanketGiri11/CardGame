package cardGame

import (
	"bufio"
	"fmt"
	"os"
)

// shuffle deck
// give each player 3 cards

type TeenPati struct {
	Deck    Deck
	Field   Deck
	Players []Player
}

// InitGame : create new game instance
func (c *TeenPati) InitGame(playerName string, numberOfPlayers int) {
	fmt.Printf("\n\n------- Register Player -------\n")
	c.Players = Register([]string{playerName}, numberOfPlayers)
	// init Deck , fill deck with cards
	fmt.Printf("\n\n------- Preparing Deck -------\n")
	c.Deck = InitDeck()
	// field for playing card
	c.Field = make([]Card, 0)
}

// StartGame : start game instance
func (c *TeenPati) StartGame() {

	var isPlaying = true

	fmt.Printf("\n\n------- Shuffle Deck -------\n")
	c.Deck.Shuffle()

	fmt.Println("\n\n-------FIELD-------")
	c.Field.AddCard(c.Deck.Draw(1)[0])

	c.Field.Show()
	// all player draw card from deck

	// draw card from deck
	for index := range c.Players {
		// draw card and give each player 3 card
		fmt.Println("\n\n-------Draw-------")
		c.Players[index].DrawCards(&c.Deck, 3)
		if c.Players[index].AI == false {
			c.Players[index].ShowHand()
		}
	}

	var round = 0

	for isPlaying {
		// return card to deck
		for index := range c.Players {

			fmt.Println("\n\n=======================FIELD====================")
			c.Field.Show()
			fmt.Printf("\n\n====================%s Turn=======================\n", c.Players[index].Name)
			if c.Players[index].AI == false {
				c.Players[index].ShowHand()
			}

			var playerTurn = true
			for playerTurn {

				// select first card for last round winner
				if round > 0 && index == 0 {
					// throw card with smalest symbol and number
					if c.Players[index].AI == false {

						fmt.Printf("\n\n================%s You Win This Round Select Any Card=============", c.Players[index].Name)

						var cardSelect int
						var playerInput = true

						for playerInput {
							fmt.Print("Enter Card Position: ")
							fmt.Scan(&cardSelect)
							if cardSelect-1 < 0 || cardSelect-1 > (len(c.Players[index].Hand)-1) {
								println("Invalid card index")
							} else {
								c.Players[index].ThrowCards(cardSelect-1, &c.Field)
								playerInput = false
							}
						}

					} else {

						fmt.Printf("\n\n==================%s You Win This Round =================\n", c.Players[index].Name)
						c.Players[index].ThrowCards(0, &c.Field)
					}
					break
				} else {

					for cIdx := range c.Players[index].Hand {
						if c.Players[index].Hand[cIdx].symbol == c.Field[len(c.Field)-1].symbol {

							if c.Players[index].AI == false {
								var cardSelect int
								var playerInput = true
								for playerInput {
									fmt.Print("Enter Card Position: ")
									fmt.Scan(&cardSelect)
									if cardSelect-1 < 0 || cardSelect-1 > (len(c.Players[index].Hand)-1) {
										println("Invalid card index")
									} else {
										if c.Players[index].Hand[cardSelect-1].symbol == c.Field[len(c.Field)-1].symbol {
											c.Players[index].ThrowCards(cardSelect-1, &c.Field)
											playerInput = false
										} else {
											fmt.Println("Cannot use this card, use card with same symbol")
										}
									}

								}

							} else {
								c.Players[index].ThrowCards(cIdx, &c.Field)
							}

							playerTurn = false
							break
						}
					}

					// Draw a new card until player have playable card
					for playerTurn {

						if c.Players[index].AI == false {
							fmt.Print("You dont have any playable card\n")
							fmt.Print("Press 'Enter' to draw a new card")
							bufio.NewReader(os.Stdin).ReadBytes('\n')
						}
						if len(c.Deck) > 0 {

							fmt.Println("\n\n-------Draw-------")
							c.Players[index].DrawCards(&c.Deck, 1)
							fmt.Printf("\n\n-------%s Hand-------\n", c.Players[index].Name)

							if c.Players[index].AI == false {
								c.Players[index].ShowHand()
							}

							if c.Players[index].Hand[len(c.Players[index].Hand)-1].symbol == c.Field[len(c.Field)-1].symbol {
								c.Players[index].ThrowCards(len(c.Players[index].Hand)-1, &c.Field)
								playerTurn = false
								break
							}

						} else {

							// draw card from field because deck is empty
							fmt.Println("\n\n-------Penalty Draw From FIELD-------")
							c.Players[index].DrawCards(&c.Field, 1)
							fmt.Printf("\n\n-------%s Hand-------\n", c.Players[index].Name)

							if c.Players[index].AI == false {
								c.Players[index].ShowHand()
							}

							playerTurn = false
							break
						}
					}
				}

			}

			if len(c.Players[index].Hand) < 1 {
				fmt.Printf("\n\n====================%s WIN==================\n", c.Players[index].Name)
				isPlaying = false
				break
			}

		}

		//fmt.Printf("\n\n-------ROUND %d OVER-------\n", round)
		round = round + 1
		// re order player based on last round winner
		var biggestHand = c.Players[0].LastPlay
		var roundWinnerIdx = 0
		for index := range c.Players {
			if IsCardHigher(c.Players[index].LastPlay, biggestHand) {
				biggestHand = c.Players[index].LastPlay
				roundWinnerIdx = index
			}
		}

		if roundWinnerIdx > 0 {
			var tempPlayer = make(Players, len(c.Players))
			copy(tempPlayer, c.Players)
			var newCounter = 0
			for index := range c.Players {

				if (roundWinnerIdx + index) <= (len(c.Players) - 1) {
					c.Players[index] = tempPlayer[roundWinnerIdx+index]
				} else {
					c.Players[index] = tempPlayer[newCounter]
					newCounter = newCounter + 1
				}
			}

		}
	}

	fmt.Println("\n\n=============GAME OVER================")

}
