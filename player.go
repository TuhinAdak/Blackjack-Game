// player.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
	Score int
}

type Deck []Card

var suits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
var values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

func NewDeck() Deck {
	var deck Deck
	for _, suit := range suits {
		for i, value := range values {
			score := i + 2
			if value == "Jack" || value == "Queen" || value == "King" {
				score = 10
			} else if value == "Ace" {
				score = 11 // Ace can be 1 or 11, handled later
			}
			deck = append(deck, Card{Suit: suit, Value: value, Score: score})
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range *d {
		j := rand.Intn(len(*d))
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

func (d *Deck) Draw() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

type Player struct {
	Name  string
	Hand  []Card
	Score int
}

func NewPlayer(name string) *Player {
	return &Player{Name: name, Hand: []Card{}, Score: 0}
}

func (p *Player) AddCard(card Card) {
	p.Hand = append(p.Hand, card)
	p.Score += card.Score
}

func (p *Player) Reset() {
	p.Hand = nil
	p.Score = 0
}

func (p *Player) ShowHand() {
	fmt.Printf("%s's Hand: ", p.Name)
	for _, card := range p.Hand {
		fmt.Printf("%s of %s, ", card.Value, card.Suit)
	}
	fmt.Printf("Score: %d\n", p.Score)
}
