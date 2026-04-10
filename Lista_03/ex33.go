package main

import f "fmt"

func main() {
	var n1, n2 int
	Q := 0
	R := 0

	f.Println("Escreva dois numeros inteiros:")
	f.Scan(&n1, &n2)

	for n1 >= n2 {
		n1 -= n2
		Q++
		R = n1
	}
	f.Printf(" O quociente: %d resto é %d", Q, R)
}
