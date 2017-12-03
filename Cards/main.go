package main //flags this as an executable package

func main() {
	cards := newDeck()
	cards.saveToFile("myDeck.txt")
	cards2 := newDeckFromFile("myDeck.txt")
	cards2.shuffle()
	cards2.print()
}
