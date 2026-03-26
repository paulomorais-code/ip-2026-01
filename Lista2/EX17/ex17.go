package main

import "fmt"

func main() {
	var conta int
	var tipo string
	var consumo float64
	var valorTotal float64

	fmt.Print("Digite o número da conta: ")
	fmt.Scan(&conta)

	fmt.Print("Digite o tipo de consumidor (R - Residencial, C - Comercial, I - Industrial): ")
	fmt.Scan(&tipo)

	fmt.Print("Digite o consumo em m³: ")
	fmt.Scan(&consumo)

	// Lógica de cálculo baseada no tipo
	switch tipo {
	case "R", "r": // Residencial
		valorTotal = 5.00 + (consumo * 0.05)

	case "C", "c": // Comercial
		valorTotal = 500.00
		if consumo > 80 {
			valorTotal += (consumo - 80) * 0.25
		}
	}
	fmt.Printf("Conta: %d\n", conta)
	fmt.Printf("Valor a pagar: R$ %.2f\n", valorTotal)
}
