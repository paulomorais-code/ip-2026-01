package main

import "fmt"

func exponencial(x, n int) int {
	if n == 1 {
		return x
	}
	return x * exponencial(x, n-1)
}

func main() {
	var x, n int
	fmt.Println("Informe a base e o expoente:")
	fmt.Scan(&x, &n)
	fmt.Println(exponencial(x, n))

}
