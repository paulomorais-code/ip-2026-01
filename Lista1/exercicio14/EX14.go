package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a float64

	fmt.Scan(&h, &a)

	Ab := 3 * math.Pow(a, 2) * (math.Sqrt(3) / 2)
	V := (1.0 / 3.0) * Ab * h

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", V)
}
