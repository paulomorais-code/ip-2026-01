package main

import "fmt"

func main() {
	var salario float64

	fmt.Scan(&salario)

	if salario <= 300 {
		salario = salario * 1.50
	} else {
		salario = salario * 1.30
	}
	fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salario)
}
