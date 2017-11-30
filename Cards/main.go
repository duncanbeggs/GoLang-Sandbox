package main //flags this as an executable package

func main() {
	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
