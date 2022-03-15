package main

func main() {
	// a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	// a = append(a, 56)
	// fmt.Println(a)
	// fmt.Println(len(a))

	// darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// dslice := darr[2:5]
	// fmt.Println("array before", darr)
	// for i := range dslice {
	// 	dslice[i]++
	// }
	// fmt.Println("slice after", dslice)
	// // changes to the slice are reflected in the array
	// fmt.Println("array after", darr)
	// // When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.

	// The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d and \ncapacity %d\n", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6

	// Regarding Capacity of slices on addition
	// cars := []string{"Ferrari", "Honda", "Ford"}
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	// cars = append(cars, "Toyota")
	// fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	// // Regarding ... operator to append slices on another slice
	// cars := []string{"Ferrari", "Honda", "Ford"}
	// bikes := []string{"Activa", "Jupiter", "Wagon"}
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	// cars = append(cars, bikes...)
	// fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
}
