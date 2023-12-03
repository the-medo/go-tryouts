package pointers

import "fmt"

type Coordinate2 struct {
	x      int
	y      int
	values *[]int
}

func RunExamples3() {
	// ================= int
	coord := Coordinate2{
		x:      10,
		y:      10,
		values: &[]int{1, 2, 3},
	}

	fmt.Println("Original: ", coord)
	modifyCoord2NoPointer(coord)
	fmt.Println("Modify value no pointer: ", coord, *coord.values)

	modifyCoord2Pointer(&coord)
	fmt.Println("Modify value pointer: ", coord, *coord.values)
}

func modifyCoord2NoPointer(c Coordinate2) {
	c.x = 20
	c.y = 100
	*c.values = append(*c.values, 200)
}

func modifyCoord2Pointer(c *Coordinate2) {
	c.x = 20
	c.y = 100
	*c.values = append(*c.values, 200)
}
