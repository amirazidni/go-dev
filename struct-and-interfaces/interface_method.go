// Interface Methods

package main

import (
	"fmt"
	"math"
)

// Interface
type geometry interface {
	area() float64
	perim() float64
}

// Objects
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

type line struct {
	length float64 // This object has no area or perim method!
}

// Methods
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Function
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Println("Rectangle measurement:")
	measure(r)
	fmt.Println("Circle measurement:")
	measure(c)

	//    l := line{length: 7}
	//    measure(l)
}
