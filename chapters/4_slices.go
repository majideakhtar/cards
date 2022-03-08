package chapters

import "fmt"

func main() {
	cards := []string{"Ace of Spades", "Five of Diamonds"}
	// fmt.Println(cards)

	cards = append(cards, "Six of Hearts")
	fmt.Println(cards)

	// for i, card := range cards {
	// 	fmt.Printf("\nIndex is %v, Value is %v\n", i, card)
	// }

	for _, card := range cards {
		fmt.Printf(" Value is %v\n", card)
	}
}
