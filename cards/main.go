package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards.txt")

	cardsFromFile := newDeckFromFile("cards.txt")
	cardsFromFile.print()
}
