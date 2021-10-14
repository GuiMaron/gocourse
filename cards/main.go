package main

import (
	"fmt"
	"os"
)

func main() {

	deck := newDeck()
	deck.shuffle()

	deck.saveToFile("deck.txt")
	cards, e := readFromFile("deck.txt")

	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}

	hand, _ := deal(cards, 5)

	hand.print()

}
