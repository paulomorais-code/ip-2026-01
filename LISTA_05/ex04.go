package main

import f "fmt"

func main() {
	var lista [10]int
	contagem := make(map[int]int)

	for i := 0; i < 10; i++ {
		f.Printf("Escreva o %d ° numero\n", i+1)
		f.Scan(&lista[i])

		contagem[lista[i]]++
	}

	for numero, qtd := range contagem {
		if qtd > 1 {
			f.Printf("O número %d aparece %d vezes\n", numero, qtd)
		}
	}
}
