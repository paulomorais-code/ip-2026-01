package main

import (
	"fmt"
)

func inverter(l []int) {
	n := len(l) - 1
	if n <= 0 {
		return
	}
	l[0], l[n] = l[n], l[0]
	inverter(l[1:n])
}

func main() {
	lista := []int{1, 2, 3, 4, 5, 6, 7, 8}

	inverter(lista)
	fmt.Println(lista)
}