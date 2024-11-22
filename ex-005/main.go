package main

import "fmt"

// Method associated with the Player struct
func (p Player) players() {
	fmt.Println("Welcome to the team", p.fname)
}

func main() {
	// Initialize Player struct with correct field names
	p1 := Player{
		id:     1,
		fname:  "john michael",
		lname:  "echavez",
		bday:   "09-16-1990",
		sports: "basketball",
		status: 1,
	}
	p1.players()
}
