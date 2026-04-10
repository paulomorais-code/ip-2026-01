package main

import f "fmt"

// funcao que verifica se o numero e primo
func ePrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n1, n2 int

	f.Print("Digite o valor de N1: ")
	f.Scan(&n1)
	f.Print("Digite o valor de N2: ")
	f.Scan(&n2)

	inicio, fim := n1, n2
	if n1 > n2 {
		inicio, fim = n2, n1
	}

	f.Printf("Números primos entre %d e %d:\n", inicio, fim)
	for i := inicio; i <= fim; i++ {
		if ePrimo(i) {
			f.Printf("%d ", i)
		}
	}
}
