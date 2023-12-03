package pointers

import "fmt"

type Coordinate struct {
	x      int
	y      int
	values []int
}

func RunExamples2() {
	// ================= int
	coord := Coordinate{
		x:      10,
		y:      10,
		values: []int{1, 2, 3},
	}

	fmt.Println("Original: ", coord)
	modifyCoordNoPointer(coord)
	fmt.Println("Modify value no pointer: ", coord)

	modifyCoordPointer(&coord)
	fmt.Println("Modify value pointer: ", coord)
}

func modifyCoordNoPointer(c Coordinate) {
	c.x = 20
	c.y = 100
	c.values = append(c.values, 200)
}

func modifyCoordPointer(c *Coordinate) {
	c.x = 20
	c.y = 100
	c.values = append(c.values, 200)
}
