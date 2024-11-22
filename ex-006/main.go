package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.length * s.width
}

func (s Circle) Area() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

func info(s Shape) float64 {
	return s.Area()
}

func main() {
	s1 := Square{
		length: 10,
		width:  10,
	}
	c1 := Circle{
		radius: 15,
	}
	fmt.Println(info(c1))
	fmt.Println(info(s1))
}
