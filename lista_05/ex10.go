package main

import f "fmt"

func main() {
	var lista [50]uint64

	lista[0] = 1
	lista[1] = 1

	for i := 2; i < 49; i++ {
		lista[i] = lista[i-1] + lista[i-2]
	}
	f.Println(lista)

}
