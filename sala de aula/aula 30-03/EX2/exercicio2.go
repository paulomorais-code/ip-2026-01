package main

import "fmt"

func main() {
	const numNotas int = 5
	var nota [numNotas]float64
	var soma float64 = 0

	for i := 0; i < numNotas; i++ {
		fmt.Printf("Digite nota %d:", i+1)
		fmt.Scan(&nota[i])
		soma += nota[i]
	}
	fmt.Printf("A soma das notas é: %.2f\n", soma)
}
