package main

import "fmt"

type deck []string

func main() {
	cards := newDeck()
	fmt.Println(cards)

}
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)

		}
	}
	return cards
}
