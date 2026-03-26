package main

import "fmt"

func main() {
	var a int
	fmt.Println("escreva um numero inteiro:")
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println(a, "é positivo")
	} else {
		if a < 0 {
			fmt.Println(a, "é negativo")
		} else {
			fmt.Println(a, "é nulo")
		}
	}
}
