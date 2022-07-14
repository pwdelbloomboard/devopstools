package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	fmt.Println("activating function area() with input r which is a rect struct")
	fmt.Println("in func area() multiplying r.width * r.height")
	return r.width * r.height
}
func (r rect) perim() float64 {
	fmt.Println("activating function perim() with input r which is a rect struct")
	fmt.Println("in func perim() multiplying 2*r.width + 2*r.height")
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
