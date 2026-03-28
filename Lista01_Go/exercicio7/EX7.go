package main

import "fmt"

func main() {
	var f, polegadas float64

	fmt.Scan(&f)
	fmt.Scan(&polegadas)

	c := (5*f - 160) / 9
	mm := polegadas * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", c)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", mm)
}
