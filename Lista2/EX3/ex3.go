package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("escreva dois numeros:")
	fmt.Scan(&a, &b)

	if (a + b) > 20 {
		valor := a + b + 8
		fmt.Println(valor)
	} else {
		valor := a + b - 5
		fmt.Println(valor)
	}
}
