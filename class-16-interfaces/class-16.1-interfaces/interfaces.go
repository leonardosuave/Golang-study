package main

import (
	"fmt"
	"math"
)

//Interface params from calculate function
type shape interface {
	area() float64
}

func calculateArea(dimensions shape) {
	fmt.Println(dimensions.area())
}

// Add method to each possible struct used in calculateArea() functions params.
// The params is shape interface. So the params need to have each sign from interface, like area().
type square struct {
	height float64
	width  float64
}
func (v square) area() float64 {
	return v.height * v.width
}

type circle struct {
	radius float64
}
func (v circle) area() float64 {
	return math.Pi * math.Pow(v.radius, 2)
}

func main() {
	squareDimension := square{10, 5}
	calculateArea(squareDimension)

	circleDimension := circle{10}
	calculateArea(circleDimension)
}