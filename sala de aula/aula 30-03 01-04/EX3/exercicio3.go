package main

import "fmt"

func main() {
	const numNotas int = 10
	var valor [numNotas]float64
	var cont int = 10

	for i := 0; i < numNotas; i++ {
		fmt.Printf("Digite nota %d:", i+1)
		fmt.Scan(&valor[i])
	}
	for cont > 0 {
		fmt.Printf("%v ", valor[cont-1])
		cont = cont - 1
	}
}
