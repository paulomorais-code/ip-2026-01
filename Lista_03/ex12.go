package main

import (
	f "fmt"
)

func main() {
	var x float64
	f.Print("Digite o valor de X: ")
	f.Scan(&x)

	soma := x
	fatorial := 1.0
	sinal := -1.0

	for i := 1; i <= 20; i++ {
		fatorial *= float64(i)
		termo := x / fatorial

		soma += sinal * termo
		sinal *= -1
	}

	f.Println("Resultado da soma:", soma)
}
