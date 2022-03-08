package chapters

import (
	"fmt"
)

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func main() {
	area, perimeter := rectProps(130.7, 1.5)
	fmt.Printf("Area %f \nPerimeter %f\n", area, perimeter)
}
