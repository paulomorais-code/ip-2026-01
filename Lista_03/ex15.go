package main

import (
	"fmt"
	f "fmt"
)

func main() {
	var n int

	f.Println("Escreva um numero")
	f.Scan(&n)

	fmt.Print("Série: ")
	for i := 1; i <= n; i++ {
		quadrado := i * i
		fmt.Printf("%d ", quadrado)
	}
}
