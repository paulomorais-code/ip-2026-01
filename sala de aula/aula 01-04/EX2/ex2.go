package main

import "fmt"

func maior(a float64, b float64, c float64) float64 {

	var Maior float64

	if a > b && a > c {
		Maior = a
	} else if b > a && b > c {
		Maior = b
	} else {
		Maior = c
	}
	return Maior
}

func main() {
	var a, b, c float64

	fmt.Println("Escreva 3 números:")

	fmt.Scan(&a, &b, &c)

	maior(a, b, c)
	fmt.Println(maior(a, b, c))
}
