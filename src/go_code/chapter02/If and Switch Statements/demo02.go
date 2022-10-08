package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 0
	if guess < 1 || guess > 100 { // ||或者  &&并且
		fmt.Println("The guess must be between 1 and 100")
	} else { // ||或者  &&并且
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("ok")
		}
	}

}
