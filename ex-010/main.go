package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	age   int
}

func main() {
	p1 := person{
		First: "Jm",
		Last:  "Echavez",
		age:   34,
	}
	p2 := person{
		First: "John Michael",
		Last:  "Echavez",
		age:   34,
	}

	persons := []person{
		p1,
		p2,
	}

	fmt.Println(p1, p2)
	fmt.Println(persons)

	pr, _ := json.Marshal(persons)
	fmt.Println(string(pr))
}
