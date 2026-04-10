package main

import (
	f "fmt"
	"math"
)

func main() {
	x := 1.0
	soma := 0.0

	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			soma += 1 / (x * x * x)
		} else {
			soma -= 1 / (x * x * x)
		}
		x += 2
	}
	pi := math.Cbrt(soma * 32)
	f.Println("O valor de Pi é:", pi)
}
