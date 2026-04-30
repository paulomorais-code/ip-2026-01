package main

import f "fmt"

func main() {
	var a float64
	var lista []float64
	var posição []int
	cont := 0

	for i := 0; i < 10; i++ {
		f.Scan(&a)
		lista = append(lista, a)
	}
	for j := 0; j < 10; j++ {
		if lista[j] > 50 {
			cont++
			posição = append(posição, j)
		}
	}
	if cont == 0 {
		f.Println("Não há numeros maior que 50")
	} else {
		f.Printf("Há %d numero(s) maior que 50 na(s) posição(ões):%d\n", cont, posição)
	}

}
