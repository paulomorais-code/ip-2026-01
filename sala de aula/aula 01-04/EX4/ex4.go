package main

import "fmt"

func fatorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * fatorial(n-1)
}

func main() {
	var n int
	fmt.Println("Escreva um número:")
	fmt.Scan(&n)
	fmt.Println(n, "! é", fatorial(n))
}
