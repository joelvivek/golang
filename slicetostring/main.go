package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

	cards.print()

	hand, remainingCards := deal(cards, 10)
	hand.print()
	remainingCards.print()

}
