package main

import (
	"fmt"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 48
	b.CanFly = false
	fmt.Println(b)
}
