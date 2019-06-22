package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println("card is", card)

	cardnum := newCardNumber()
	fmt.Println("cardnumber is", cardnum)
}

func newCard() string {
	return "Five of Diamonds"
}

func newCardNumber() int {
	return 12
}
