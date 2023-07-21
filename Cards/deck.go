package main

import (
	"fmt"
	"io/ioutil"
	"strings"

)

// create a new type of 'deck'
// which is a slice of strings
type deck []string 

func newDeck () deck {  // anytime someone calls newDeck ,they're gonna return a value of type deck.
	cards:= deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues:= []string{"Ace","Two", "Three", "Four",}

	for _, suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards , value+ " of " +suit)
		}
	}
	return cards

}

func (d deck) print() {  // receiver argument
	for i , card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck , deck){

	return d[:handSize],d[handSize:]


}

func (d deck) toString() string { //receiver type deck
	return strings.Join ([]string(d) , ",")
	//[]string(d) type conversion : deck become slice of string
	// Strings

	
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
 