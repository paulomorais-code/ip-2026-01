package main

import f "fmt"

func main() {
	var n int
	f.Println("Informe um valor:")
	f.Scan(&n)

	if n <= 0 || n >= 335 {
		return
	}
	valor := 1000.0
	soma := 1000.0

	for i := 2; i <= n; i++ {
		valor = valor - 3
		if i%2 == 0 {
			soma -= valor / float64(i)
		} else {
			soma += valor / float64(i)
		}

	}
	f.Printf("A soma é %.2f\n", soma)
}
