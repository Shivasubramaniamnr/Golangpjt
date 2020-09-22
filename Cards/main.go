package main

func main() {
	cards := []string{"Ace of Spades", "Two of Clover", newCard()}

	println(cards)
}

func newCard() string {
	return "five of diamonds"
}
