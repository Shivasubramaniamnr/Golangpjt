package main

import "fmt"

func main() {
	card := NewCard()

	fmt.Println(card)
}

// newcard function to create new card deck
func NewCard() string {
	return "five of diamonds"
}
