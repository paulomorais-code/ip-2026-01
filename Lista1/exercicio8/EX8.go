package main

import "fmt"

func main() {
	var Altura, raio float64
	fmt.Scan(&raio)
	fmt.Scan(&Altura)
	pi := 3.14159

	AreaT := (2 * pi * Altura * raio) + 2*(pi*raio*raio)
	custo := AreaT * 100

	fmt.Printf("O VALOR DO CUSTO E: %.2f\n", custo)
}
