package main

import "fmt"

func main() {
	str1 := "Ace of Spades"
	str2 := "Five of Diamonds"
	var num int
	card1, card2, num := newCard(str1, str2)

	fmt.Println(card1)
	fmt.Println(card2)
	fmt.Println(num)
}

func newCard(str1 string, str2 string) (string, string, int) {
	return str1, str2, 10
}

// // Variable Declarations
// func main() {
// 	// var card string = "Ace of Spades"
// 	// var card = "Ace of Spades"
// 	card := "Ace of Spades"
// 	fmt.Println(card)

// 	card = "Five of Diamonds"
// 	fmt.Println(card)
// }
