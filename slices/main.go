package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamond", "Two of Spades"}

	cards = append(cards, "Three of Spades")

	fmt.Println("Slice of Cards: ", cards)
	fmt.Println("iteration of cards")

	// iterating cards in range taking each card through loops
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

func newCardNumber() int {
	return 12

}
