package main

import (
	"fmt"
)

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	area, perimeter := rectProps(130.7, 1.5)
	fmt.Printf("Area %f \nPerimeter %f\n", area, perimeter)
}
