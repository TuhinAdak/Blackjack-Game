// game.go
package main

import (
	"fmt"
)

func playerTurn(player *Player, deck *Deck) {
	for {
		var action string
		fmt.Print("Do you want to (h)it or (s)tand? ")
		fmt.Scanln(&action)

		if action == "h" {
			player.AddCard(deck.Draw())
			player.ShowHand()

			if player.Score > 21 {
				fmt.Println("You bust! Dealer wins.")
				return
			}
		} else if action == "s" {
			break
		} else {
			fmt.Println("Invalid action. Please enter 'h' or 's'.")
		}
	}
}

func dealerTurn(dealer *Player, deck *Deck) {
	fmt.Println("Dealer's turn:")
	dealer.ShowHand()
	for dealer.Score < 17 {
		dealer.AddCard(deck.Draw())
		dealer.ShowHand()
	}
}

func determineWinner(player *Player, dealer *Player) {
	if dealer.Score > 21 {
		fmt.Println("Dealer busts! You win!")
	} else if player.Score > dealer.Score {
		fmt.Println("You win!")
	} else if player.Score < dealer.Score {
		fmt.Println("Dealer wins!")
	} else {
		fmt.Println("It's a tie!")
	}
}
