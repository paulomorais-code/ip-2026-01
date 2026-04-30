package main

import f "fmt"

func main() {
	var lista []int

	for i := 100; i > 0; i-- {
		lista = append(lista, i)
	}
	f.Println(lista)
}
