package main

import "fmt"

func main() {
	var conta int
	var consumo float64
	var tipo string
	var valor float64

	fmt.Scan(&conta, &consumo, &tipo)
	if tipo == "R" {
		valor = 5 + (consumo * 0.05)
	} else if tipo == "C" {
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + (consumo-80)*0.25
		}
	} else if tipo == "I" {
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + (consumo-100)*0.04
		}
	}
	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}
