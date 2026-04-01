package main

import "fmt"

func main() {
	var a int
	fmt.Println("Digite um número de 3 algarismo:")
	fmt.Scan(&a)

	if a <= 99 || a > 999 {
		fmt.Println(a, "Não possui 3 algarismos")
		return
	}
	b := (a / 10) % 10

	fmt.Print(b, " é o algarismo da dezena")
}
