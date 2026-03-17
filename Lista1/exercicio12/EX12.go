package main

import "fmt"

func main() {
	var horas int
	var total float64

	fmt.Scan(&horas)

	blocos := horas / 3
	resto := horas % 3

	total = float64(blocos)*10 + float64(resto)*5

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", total)
}
