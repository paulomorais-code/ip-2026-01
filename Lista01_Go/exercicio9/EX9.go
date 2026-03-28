package main

import "fmt"

func main() {
	var A, B, C float64
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	delta := B*B - (4 * A * C)

	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}
