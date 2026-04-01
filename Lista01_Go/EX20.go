package main

import "fmt"

func main() {
	var hr, min, seg, segt int
	fmt.Scan(&hr)
	fmt.Scan(&min)
	fmt.Scan(&seg)

	segt = (hr * 3600) + (min * 60) + seg

	fmt.Printf("O VALOR EM SEGUNDOS E = %d\n", segt)
}
