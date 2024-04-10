package main

func main() {
	cards := newDeck()
	hand, remaningCards := deal(cards, 1)

	hand.print()
	remaningCards.print()
}
