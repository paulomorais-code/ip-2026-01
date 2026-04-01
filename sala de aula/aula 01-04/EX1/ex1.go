package main

import "fmt"

func soma(a float64, b float64) {
	somaN := a + b
	fmt.Print("soma = ", somaN)
}

func main() {
	var a, b float64
	fmt.Println("Digite dois números: ")
	fmt.Scan(&a, &b)
	soma(a, b)
}
