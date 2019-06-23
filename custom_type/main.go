package main

func main() {
	cards := newDeck()

	cards.print()

	hand, remainingCards := deal(cards, 10)
	hand.print()
	remainingCards.print()

}
