package main

import f "fmt"

func main() {
	var lista []int
	var a int
	soma := 0

	for i := 0; i < 100; i++ {
		f.Printf("Escreva o %d° número\n", i+1)
		f.Scan(&a)
		lista = append(lista, a)
	}
	j := 99
	for i := 0; i < 50; i++ {
		soma += (lista[i] + lista[j]) * (lista[i] + lista[j]) * (lista[i] + lista[j])
		j--
	}
	f.Println(soma)

}
