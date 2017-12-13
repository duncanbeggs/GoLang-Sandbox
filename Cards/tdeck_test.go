//how do we know what to test?
//it is up to you as a developer to determine
//what you really care about
//so for the newDeck() function something we care about is
package main

import (
	"os"
	"testing"
)

//1. check the length of the deck that get's returned (52)
//2. check that the very first card in the deck is the Ace of Spades
//3. check that the last card is equal to the king of clubs
//Note that the function is capitalized at the beginning
func TestNewDeck(t *testing.T) { //note that t is our 'test handler' variable
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected fird card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card king of clubs but got %v", d[len(d)-1])
	}
}

//the type *testing.T involves one of the more tricky topics around go
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("expected 52 cards in deck, got %v", len(loadedDeck))
	}

	//	os.Remove("_decktesting")
}
