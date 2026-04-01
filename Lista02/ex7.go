package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("escreva três numeros distintos")
	fmt.Scan(&a, &b, &c)
	if a == b || a == c || c == b {
		fmt.Println("não são distintos")
		return
	}
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	Maior := c
	Menor := a
	Inter := b
	fmt.Println(Menor, Inter, Maior)

}
