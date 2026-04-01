package main

import "fmt"

func main() {
	var a int
	fmt.Println("Escreva um número:")
	fmt.Scan(&a)
	if a%2 == 0 {
		fmt.Println(a, "é par")
	} else {
		fmt.Println(a, "é ímpar")
	}
}
