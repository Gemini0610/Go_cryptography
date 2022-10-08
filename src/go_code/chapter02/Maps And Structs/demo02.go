package main

import (
	"fmt"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string //a slice of string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon",
		companions: []string{
			"Liz",
			"Sam",
		},
	}
	fmt.Println(aDoctor)
}
