package main

import "fmt"

func main() {
	cards := []string{"Ace of Spades", "Two of Clover", newCard()}

	cards = append(cards, "three of spades")

	for i, card := range cards {
		fmt.Println(i, card)

	}

}

func newCard() string {
	return "five of diamonds"
}
