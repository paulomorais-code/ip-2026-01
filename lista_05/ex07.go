package main

import f "fmt"

func main() {
	var lista []int

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
		} else {
			lista = append(lista, i)
		}
	}
	f.Println(lista)
}
