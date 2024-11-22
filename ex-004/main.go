package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

// variadic parameter
func foo(x ...int) int {
	t := 0
	for _, v := range x {
		t += v
	}

	return t
}

// same with variadic w/parameter
func bar(x []int) int {
	t := 0
	for _, v := range x {
		t += v
	}

	return t
}
