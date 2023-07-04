package main

import (
	"fmt"
	"math"
)

type square struct {
	length    float64
	width     float64
	mathShape string
}

type circle struct {
	radius    float64
	mathShape string
}

func (sqr square) area() (string, float64) {
	return sqr.mathShape, sqr.length * sqr.width
}

func (cir circle) area() (string, float64) {
	return cir.mathShape, math.Pi * cir.radius * cir.radius
}

// Now both type circle and square have method area()
// We'll implement an interface of those type with method area() in the interface

type shape interface {
	area() (string, float64)
}

func info(sh shape) {
	objectShape, ft2 := sh.area()
	fmt.Printf("The area of %v is %v\n", objectShape, ft2)
}
func main() {
	mysquare := square{
		length:    2,
		width:     5,
		mathShape: "a square",
	}
	mycircle := circle{
		radius:    3,
		mathShape: "a circle",
	}
	info(mysquare)
	info(mycircle)
}
