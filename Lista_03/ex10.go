package main

import f "fmt"

func main() {
	var num int

	f.Println("Informe um numero maior ou igual a 3:")
	f.Scan(&num)
	a, b := 0, 1

	f.Printf("Sequência de Fibonacci (%d termos):\n", num)
	f.Printf("%d – %d", a, b)

	// Cálculo dos termos subsequentes
	for i := 3; i <= num; i++ {
		proximo := a + b
		f.Printf(" – %d", proximo)

		// Atualização das variáveis para o próximo ciclo
		a = b
		b = proximo
	}
	f.Println()
}
