package main

import f "fmt"

func main() {
	var x int

	f.Println("Digite um valor inteiro")
	f.Scan(&x)
	soma := 0

	for i := 1; i <= x; i++ {
		soma += i
	}
	f.Println("A soma é ", soma)
}
