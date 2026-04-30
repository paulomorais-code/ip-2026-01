package main

import f "fmt"

func main() {
	var lista [20]int
	contagem := make(map[int]int)

	for i := 0; i < 20; i++ {
		f.Println("Escreva o número sorteado:")
		f.Scan(&lista[i])
		if lista[i] < 1 || lista[i] > 6 {
			return
		}

		contagem[lista[i]]++
	}
	for i := 1; i <= 6; i++ {
		f.Printf("O número %d aparece %d vezes\n", i, contagem[i])
	}
}
