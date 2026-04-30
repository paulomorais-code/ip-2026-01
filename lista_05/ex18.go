package main

import f "fmt"

func main() {
	var lista [10]int
	var a int

	for i := 0; i < 10; i++ {
		f.Println("Escreva um valor:")
		f.Scan(&a)

		j := i - 1
		for j >= 0 && lista[j] > a {
			lista[j+1] = lista[j]
			j--
		}
		lista[j+1] = a
	}
	f.Println(lista)
}
