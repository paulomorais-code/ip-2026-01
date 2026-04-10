package main

import f "fmt"

func main() {
	var grão uint64 = 1
	var soma uint64 = 0

	for i := 0; i < 64; i++ {
		soma += grão
		grão *= 2
	}
	f.Println("O numero de grãos é:", soma)
}
