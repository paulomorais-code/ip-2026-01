package main

import f "fmt"

func main() {
	var a float64
	var lista []float64

	for i := 0; i < 30; i++ {
		f.Scan(&a)
		lista = append(lista, a)
	}
	for i := 0; i < 30; i++ {
		if i%2 == 0 {
			lista[i] = lista[i] * 2
		} else {
			lista[i] = lista[i] * 3
		}
	}
	f.Println(lista)
}
