package main

func main() {
	cards := newDeck()
	hand, remamingCards := deal(cards, 5)

	hand.print()
	remamingCards.print()
}
