package main

import "fmt"

func main() {
	var dia int
	var categoria string
	var valor float64

	fmt.Println("Digite o dia da semana:")
	fmt.Scan(&dia)

	fmt.Println("Escreva o tipo de cd: normal ou lançamento: ")
	fmt.Scan(&categoria)

	fmt.Println("Escreva o valor do cd")
	fmt.Scan(&valor)

	switch dia {
	case 1, 4, 6, 7:
		switch categoria {
		case "normal", "Normal":
			fmt.Printf("preço é R$: %.2f\n", valor)
		case "lançamento", "Lançamento":
			valor = 1.15 * valor
			fmt.Printf("o preço é R$: %.2f\n", valor)
		default:
			fmt.Println("Não é válido")
		}
	case 2, 3, 5:
		valor = valor * 0.6
		fmt.Printf("o valor é R$: %.2f\n", valor)

	}
}
