package main

import f "fmt"

func main() {
	var lista [10]int

	for i := 0; i < 10; i++ {
		f.Printf("Escreva o %d ° numero\n", i+1)
		f.Scan(&lista[i])
	}

	menor := lista[0]
	pos := 0

	for i := 1; i < 10; i++ {
		if lista[i] < menor {
			menor = lista[i]
			pos = i
		}
	}
	f.Printf("o menor número é %d e sua posição é %d\n", menor, pos)
}
