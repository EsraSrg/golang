package main

func main(){
	//var card string = "Ace of Spades" equals this:
	//cards := deck{"Ace of Diamonds", newCard()}
	cards:=newDeck()
	

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	

}


