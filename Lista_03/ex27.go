package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Digite x: ")
	fmt.Scan(&x)

	cosseno := 1.0
	termo := 1.0

	for i := 1; i < 20; i++ {
		termo *= -x * x / float64((2*i-1)*(2*i))
		cosseno += termo
	}

	fmt.Printf("cos(%.2f) = %.6f\n", x, cosseno)
}
