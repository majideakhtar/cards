package main

func main() {
	// cards := newDeckFromFile("my_cards")
	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
	// handSize := 5
	// hand, remainingCards := deal(cards, handSize)

	// // cards.print()
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	// my_cards :=

	// For deck print()
	// for i, card := range cards {
	// 	fmt.Printf("\nIndex is %v, Value is %v\n", i, card)
	// }

	// for _, card := range cards {
	// 	fmt.Printf(" Value is %v\n", card)
	// }
}

// func main() {
// 	str1 := "Ace of Spades"
// 	str2 := "Five of Diamonds"
// 	var num int
// 	card1, card2, num := newCard(str1, str2)

// 	fmt.Println(card1)
// 	fmt.Println(card2)
// 	fmt.Println(num)
// }

// func newCard(str1 string, str2 string) (string, string, int) {
// 	return str1, str2, 10
// }

// // Variable Declarations
// func main() {
// 	// var card string = "Ace of Spades"
// 	// var card = "Ace of Spades"
// 	card := "Ace of Spades"
// 	fmt.Println(card)

// 	card = "Five of Diamonds"
// 	fmt.Println(card)
// }
