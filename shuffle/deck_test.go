package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 2006 {
		t.Errorf("Expeted deck of length 16, but got %v ", len(d))
	}
}

func TestSaveToDeckAndNewDeckTesFile(t *testing.T) {

}