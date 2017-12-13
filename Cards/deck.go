pckage main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os" //These are platform indepent functions! Works equally well on Windows vs. Linux, etc.
	"strings"
	"time"
)

//Create a new type of deck which is a slice of strings
//this is sort of like a subclass that extends []string
type deck []string //a deck of cards

//will return a value of type 'deck'
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
		"Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	deckString := strings.Join([]string(d), ", ")
	return deckString
}

//this function will take a byte array and turn it into our deck type
func toDeck(fileContents []byte) deck {
	deckString := string(fileContents)
	s := strings.Split(deckString, ", ")
	return deck(s)
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	fileBytes, readError := ioutil.ReadFile(fileName)
	if readError != nil { //error handling
		fmt.Println("Error: ", readError)
		os.Exit(1)
	}
	return toDeck(fileBytes) //before returning, use toDeck() to convert byte array to
}

func (d deck) shuffle() {
	//generate random number generator using nanosecond times
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := range d {
		//Intn(n) is a very useful function for random number generation
		//will give you a number between 0 and n
		//also useful is the len(slice) function
		np := r.Intn(len(d) - 1)
		//fancy one-line syntax for switching d[i] with d[np]
		d[i], d[np] = d[np], d[i]
	}
}
