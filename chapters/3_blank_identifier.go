package chapters

import (
	"fmt"
)

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
func main() {
	//discarding return value for perimeter from the below function call
	area, _ := rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("Area %f ", area)
}
