package main
import "fmt"

func main(){
	//var card string = "Ace of Spades" equals this:
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards:=newDeck()
	

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()
	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting))
	cards := newDeck()
	fmt.Println(cards.toString())

	

}


