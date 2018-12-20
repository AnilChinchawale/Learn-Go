package main

func main() {

	// Cards := deck{"Ace of Spades", newDeck()}
	// for i, card := range Cards {
	// 	fmt.Println(i, card)
	// }

	// Cards.print()
	// Cards := newDeck()
	// hand, remainingDeck := deal(Cards, 3)
	// hand.print()
	// remainingDeck.print()
	// Cards := newNumList()
	// Cards.print()
	// Cards.saveToHDD("myfile.txt")
	Cards := newDeckFromFile("myfile.txt")
	Cards.print()
}
