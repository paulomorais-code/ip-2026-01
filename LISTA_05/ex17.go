package main

import f "fmt"

func ePrimo(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var lista [10]int
	var lista1 []int

	for i := 0; i < 10; i++ {
		f.Printf("Escreva o %d ° numero: ", i+1)
		f.Scan(&lista[i])
	}
	for i := 0; i < 10; i++ {
		if ePrimo(lista[i]) {
			lista1 = append(lista1, lista[i], i)
		}
	}
	f.Println(lista1)
}
