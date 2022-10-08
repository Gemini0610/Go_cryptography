package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California": 123131546,
		"Texas":      121548976,
	}
	if pop, ok := statePopulations["New York"]; ok {
		fmt.Println(pop)
	}
}
