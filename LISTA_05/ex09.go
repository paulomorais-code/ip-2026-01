package main

import f "fmt"

func main() {
	var lista []float64
	var maior []float64
	var a float64
	var media float64
	qtd := 10.0
	soma := 0.0

	for i := 0; i < 10; i++ {
		f.Printf("Escreva a %d° altura:", i+1)
		f.Scan(&a)
		soma += a
		lista = append(lista, a)
	}
	media = soma / qtd

	for _, l := range lista {
		if l > media {
			maior = append(maior, l)
		}
	}
	f.Println("A altura dos atletas acima da média é:", maior)
}
