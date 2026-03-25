package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	valor := a*3 + b*2 + c
	fmt.Println(valor)
	fmt.Printf("%d, %d, %d\n", a, b, c)
}
