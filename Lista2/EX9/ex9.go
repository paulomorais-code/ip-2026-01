package main

import "fmt"

func main() {
	var valor float64
	fmt.Println("Escreva o valor:")
	fmt.Scan(&valor)
	if valor < 10 {
		venda := valor * 1.7
		fmt.Printf("%.2f\n", venda)
	}
	if valor >= 10 && valor < 30 {
		venda := valor * 1.5
		fmt.Printf("%.2f\n", venda)
	}
	if valor >= 30 && valor < 50 {
		venda := valor * 1.4
		fmt.Printf("%.2f\n", venda)
	}
	if valor >= 50 {
		venda := valor * 1.3
		fmt.Printf("%.2f\n", venda)
	}
}
