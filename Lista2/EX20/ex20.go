package main

import "fmt"

func main() {
	var valor float64
	var tipo int

	fmt.Println("O valor é:")
	fmt.Scan(&valor)

	fmt.Println("Qual será o meio de pagamento?")
	fmt.Scan(&tipo)

	switch tipo {
	case 1:
		valor := valor * 0.9
		fmt.Printf("o valor final será %.2f\n", valor)

	case 2:
		valor := valor * 0.95
		fmt.Printf("o valor final será %.2f\n", valor)

	case 3:
		valor := valor
		fmt.Printf("o valor final em 2x será %.2f\n", valor)

	case 4:
		valor := valor * 1.1
		fmt.Printf("o valor final em 3x será %.2f\n", valor)

	default:
		fmt.Println("Valor não compátivel")
	}

}
