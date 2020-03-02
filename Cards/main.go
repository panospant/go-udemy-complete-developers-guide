package main

func main() {
	//var card string = newCard()
	// card := newCard()
	var cards = newDeck()

	// var hand, remainingDeck = deal(cards, 5)
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()

	// fmt.Println("")
	// fmt.Println("Remaining deck is: ")
	// fmt.Println("")

	// remainingDeck.print()

	// cards.saveToFile("my_cards")

	cards = newDeckFromFile("my_cards")

	cards.shuffle()

	cards.print()
}
