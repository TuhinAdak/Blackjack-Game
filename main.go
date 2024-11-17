// main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Blackjack!")

	// Initialize deck and shuffle
	deck := NewDeck()
	deck.Shuffle()

	// Create players
	player := NewPlayer("Player")
	dealer := NewPlayer("Dealer")

	// Initial deal
	for i := 0; i < 2; i++ {
		player.AddCard(deck.Draw())
		dealer.AddCard(deck.Draw())
	}

	// Show hands
	player.ShowHand()
	fmt.Printf("Dealer's showing card: %s of %s\n", dealer.Hand[0].Value, dealer.Hand[0].Suit)

	// Player's turn
	playerTurn(player, &deck) // Pass the address of the deck

	// Dealer's turn
	dealerTurn(dealer, &deck) // Pass the address of the deck

	// Determine winner
	determineWinner(player, dealer)
}
