package main

import "fmt"

func BuscaBinaria(l []int, x int) int {
	n := len(l)
	e := 0
	d := n - 1
	for e <= d {
		m := (e + d) / 2
		if l[m] == x {
			return m
		}
		if l[m] < x {
			e = m + 1
		}
		if l[m] > x {
			e = m - 1

		}
	}
	return -1
}

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(BuscaBinaria(lista, 7))
}
