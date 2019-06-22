package main

func main() {
	cards := deck{"Ace of Diamond", "Two of Spades"}
	cards = append(cards, "Three of Spades")

	cards.print()

}
