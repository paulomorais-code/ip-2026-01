package main

import f "fmt"

func main() {
	var lista [10]int
	var lista2 [5]int

	for i := 0; i < 10; i++ {
		f.Println("Escreva um valor para lista1:")
		f.Scan(&lista[i])
	}

	for i := 0; i < 5; i++ {
		f.Println("Escreva um valor para lista2:")
		f.Scan(&lista2[i])
	}

	for a, l2 := range lista2 {
		for i := 0; i < 10; i++ {
			if lista[i]%l2 == 0 {
				f.Printf("O numero %d é divísivel por %d na posição %d\n", lista[i], l2, a)
			}
		}
	}
}
