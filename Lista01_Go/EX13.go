package main

import "fmt"

func main() {
	var nota float64
	var conceito string

	fmt.Scan(&nota)

	if nota >= 9.0 {
		conceito = "A"
	} else if nota >= 7.5 {
		conceito = "B"
	} else if nota >= 6.0 {
		conceito = "C"
	} else {
		conceito = "D"
	}

	fmt.Printf("NOTA = %.1f CONCEITO = %s\n", nota, conceito)
}
