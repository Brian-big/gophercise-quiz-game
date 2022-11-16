package main

import (
	"fmt"
	"math"
)


type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {

	c1 := circle{2}
	r1 := rectangle{4,5}
	
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
		
	}
	
}