package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	hand1, _ := deal(cards, 5)
	hand2, _ := deal(cards, 5)
	fmt.Println(hand1, hand2)
}
