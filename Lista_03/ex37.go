package main

import "fmt"

func main() {
	var octal int
	fmt.Print("Digite um número em base 8: ")
	fmt.Scan(&octal)

	decimal := 0
	potencia := 1

	for octal > 0 {
		digito := octal % 10

		decimal += digito * potencia
		potencia *= 8

		octal /= 10
	}

	fmt.Println("Decimal:", decimal)
}
