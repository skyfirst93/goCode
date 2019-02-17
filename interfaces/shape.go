package main

import (
	"fmt"
	"math"
)

type shape interface {
	area()
}

type circle struct {
	rad float64
}

type rect struct {
	l, b int
}

func (c *circle) area() {
	fmt.Println("area of circle", c.rad*math.Pi*c.rad)
}
func (r *rect) area() {
	fmt.Println("area of rectangle", r.l*r.b)
}

func main() {
	var s, s1 shape
	s = &circle{2}
	s1 = &rect{2, 4}
	s.area()
	s1.area()

}
