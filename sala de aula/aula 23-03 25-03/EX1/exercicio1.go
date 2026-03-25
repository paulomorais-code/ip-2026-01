package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("escreva os lados do triangulo:")
	fmt.Scan(&a, &b, &c)

	if a >= (b+c) || b >= (a+c) || c >= (a+b) {
		fmt.Print("valores inválidos, não é triangulo")
		return
	}

	if a == b && b == c && a == c {
		fmt.Println("o triangulo e equilatero")
	} else if a == b || b == c || c == a {
		fmt.Println("o triangulo e isoceles")
	} else {
		fmt.Println("o triangulo é escaleno")
	}
}
