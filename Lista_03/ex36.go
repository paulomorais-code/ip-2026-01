package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Hexadecimal: 0")
		return
	}

	hex := ""
	digitos := "0123456789ABCDEF"

	for n > 0 {
		resto := n % 16
		hex = string(digitos[resto]) + hex
		n = n / 16
	}

	fmt.Println("Hexadecimal:", hex)
}
