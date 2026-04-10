package main

import f "fmt"

func main() {
	var carlos, joao float64
	meses := 0

	f.Println("Qual é o salário de Carlos?")
	f.Scan(&carlos)
	joao = carlos / 3

	for joao < carlos {
		carlos = carlos * 1.02
		joao = joao * 1.05
		meses++
	}

	f.Printf("João passará Carlos em %d meses", meses)

}
