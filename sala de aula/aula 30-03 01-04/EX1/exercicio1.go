package main

import "fmt"

func main() {

	const numNotas int = 3
	var (
		nota [numNotas]float64
		soma float64 = 0
	)

	for i := 0; i < numNotas; i++ {
		fmt.Printf("Informe a nota %d: ", i+1)
		fmt.Scan(&nota[i])
		soma += nota[i]
	}
	media := soma / float64(numNotas)

	for i, v := range nota {
		fmt.Printf("Nota %d = %.2f\n", i+1, v)
	}
	fmt.Printf("A média das notas é: %.2f\n", media)
	if nota[0] > media {
		fmt.Println(nota[0], "está acima da média")
	}
	if nota[1] > media {
		fmt.Println(nota[1], "está acima da média")
	}
	if nota[2] > media {
		fmt.Println(nota[2], "está acima da média")
	}
}
