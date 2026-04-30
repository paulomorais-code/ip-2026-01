package main

import (
	f "fmt"
	"math"
)

func main() {
	var lista [15]int
	var qua []float64

	for i := 0; i < 15; i++ {
		f.Printf("Escreva o %d ° numero\n", i+1)
		f.Scan(&lista[i])
		if lista[i] < 0 {
			qua = append(qua, -1)
		} else {
			raiz := math.Sqrt(float64(lista[i]))
			qua = append(qua, raiz)
		}
	}
	f.Println(qua)
}
