package main

import "fmt"

func main() {
	var a int
	fmt.Println("Escreva um número:")
	fmt.Scan(&a)

	funcao := 8 / (2 - a)
	fmt.Println(funcao)
}
