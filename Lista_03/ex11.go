package main

import f "fmt"

func fatorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * fatorial(n-1)
}

func main() {
	var n int

	f.Println("Informe um valor inteiro positivo:")
	f.Scan(&n)

	if n <= 0 {
		f.Println("Numero inválido")
		return
	}
	f.Printf("o valor de %d! é %d", n, fatorial(n))

}
