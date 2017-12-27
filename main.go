package main

func main() {
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
	// fmt.Println(cards)
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// hand, remainingDeck := deal(cards, 5)

	// remainingDeck.print()
	// hand.print()
}
