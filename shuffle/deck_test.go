package main

<<<<<<< HEAD
import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {

=======
import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expeted deck of length 52, but got %v ", len(d))
	}

	if d[0] != "Ace of Spades" {
>>>>>>> ee58f98b0b745069584922ea3948ee33035f5ace
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
<<<<<<< HEAD

		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
=======
		t.Errorf("Expected four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTesFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, got %v", len(loadedDeck))
	}

	/* to remove the temp file */
	os.Remove("_decktesting")

>>>>>>> ee58f98b0b745069584922ea3948ee33035f5ace
}
