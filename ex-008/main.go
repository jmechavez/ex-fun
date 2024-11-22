package main

import "fmt"

func add(n int) int {
	return n + n
}

func addition(f func(int) int, x int) string {
	total := f(x)
	return fmt.Sprintf("total number is: %v", total)
}

func main() {
	fmt.Println(addition(add, 2))
}
