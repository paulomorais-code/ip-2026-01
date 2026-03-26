package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	fmt.Println("escreva um numero")
	fmt.Scan(&a)

	if a >= 0 {
		fmt.Println(math.Sqrt(a))
	} else {
		fmt.Println(a * a)
	}
}
