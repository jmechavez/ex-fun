package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	p1 := person{
		First: "Jm",
		Last:  "Echavez",
		Age:   34,
	}
	p2 := person{
		First: "John Michael",
		Last:  "Echavez",
		Age:   34,
	}

	persons := []person{
		p1,
		p2,
	}

	fmt.Println(p1, p2)
	fmt.Println(persons)

	pr, _ := json.Marshal(persons)
	fmt.Println(string(pr))

	var unM []person
	fmt.Println(unM)

	err := json.Unmarshal(pr, &unM)
	if err != nil {
		fmt.Println("Error Unmarshal")
	}
	fmt.Println(unM)

	for _, v := range unM {
		fmt.Println(v.First, v.Last)
	}
}
