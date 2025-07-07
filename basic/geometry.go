package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Width, Height float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Println("Area is :", s.Area())
}

func main() {
	c1 := Circle{5.6}
	r1 := Rectangle{7.6, 5.7}
	printArea(&c1)
	printArea(&r1)
}
