package main

import (
	"fmt"
	"math"
)

func main() {
	myNum := 0.123456789
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.01 {
		fmt.Println("same")
	} else {
		fmt.Println("different")
	}
}
