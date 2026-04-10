package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Binário: 0")
		return
	}

	binario := ""

	for n > 0 {
		resto := n % 2
		binario = fmt.Sprintf("%d", resto) + binario
		n = n / 2
	}

	fmt.Println("Binário:", binario)
}
