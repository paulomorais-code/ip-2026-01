package main

import "fmt"

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var n1, n2 int

	fmt.Print("Digite dois números positivos: ")
	fmt.Scan(&n1, &n2)

	mmc := (n1 * n2) / mdc(n1, n2)

	fmt.Println("MMC =", mmc)
}
