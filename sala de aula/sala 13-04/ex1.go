package main

import "fmt"

func BuscaSequencial(l []int, x int) int {
	n := len(l)
	for i := 0; i < n; i++ {
		if l[i] == x {
			fmt.Print(i)
		}
	}
	return fmt.Print(-1)
}

func main() {
	var lista = []int{9, 3, 8, 9}

	BuscaSequencial(lista, 9)

}
