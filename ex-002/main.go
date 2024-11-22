package main

import "fmt"

func main() {
	var a string = "test" // assign the value "test" to the variable "a" of type string
	var b int = 1         // assign the value 1 to the variable "b" of type int
	var c float32 = 1.1   // assign the value 1.1 to the variable "c" of type float32

	fmt.Printf("%v, %T \n", a, a)
	fmt.Printf("%v, %T \n", b, b)
	fmt.Printf("%v, %T \n", c, c)
}
