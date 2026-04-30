package main

import f "fmt"

func main() {
	var lista [50]int
	contagem := make(map[int]int)

	for i := 0; i < 50; i++ {
		f.Printf("Escreva a %d ° idade: ", i+1)
		f.Scan(&lista[i])

		contagem[lista[i]]++
	}

	maiorFreq := 0
	for _, v := range contagem {
		if v > maiorFreq {
			maiorFreq = v
		}
	}

	var modas []int
	for idade, qtd := range contagem {
		if qtd == maiorFreq {
			modas = append(modas, idade)
		}
	}

	f.Printf("Moda(s): %v\n", modas)
}
