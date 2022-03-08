package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

// type Size uint8

// const (
// 	small = Size(iota)
// 	medium
// 	large
// 	extraLarge
// )

// func main() {
// 	var m Size = 3
// 	m.toString()
// }
// func (s Size) toString() {
// 	switch s {
// 	case small:
// 		fmt.Println("Small")
// 	case medium:
// 		fmt.Println("Medium")
// 	case large:
// 		fmt.Println("Large")
// 	case extraLarge:
// 		fmt.Println("Extra Large")
// 	default:
// 		fmt.Println("Invalid Size entry")
// 	}
// }

// const (
// 	small Size = iota
// 	medium
// 	large
// 	extraLarge
// )

// func main() {
// 	fmt.Println(small)
// 	fmt.Println(medium)
// 	fmt.Println(large)
// 	fmt.Println(extraLarge)
// }
