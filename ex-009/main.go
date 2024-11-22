package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently your are bound to succeed."
	r := "The meaning of life is..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("a: %v \t%T\n", a, a)
	fmt.Printf("p: %v \t%T\n", b, b)
	fmt.Printf("a: %v \t%T\n", c, c)
	fmt.Printf("a: %v \t%T\n", d, d)
}
