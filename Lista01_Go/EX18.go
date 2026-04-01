package main

import "fmt"

func main() {
	var a1, r, n int
	var soma int = 0

	fmt.Scan(&a1, &r, &n)

	for i := 0; i < n; i++ {
		termo := a1 + i*r
		soma += termo
	}

	fmt.Println(soma)
}
