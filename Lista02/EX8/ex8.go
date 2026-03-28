package main

import "fmt"

func main() {
	var a int
	fmt.Println("Escreva um número:")
	fmt.Scan(&a)
	if a > 20 && a < 90 {
		fmt.Println(a, "está entre 20 e 90")
	} else {
		fmt.Println(a, "não está entre 20 e 90")
	}
}
