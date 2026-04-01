package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var raio, altura float64
	fmt.Println("Qual a figura escolhida? 1 para cone, 2 para cilindro e 3 para esfera")
	fmt.Scan(&a)
	fmt.Println("Qual é o valor do raio?")
	fmt.Scan(&raio)
	fmt.Println("Qual é o valor da altura?")
	fmt.Scan(&altura)

	switch a {
	case 1:
		volume := (math.Pi * math.Pow(raio, 2) * altura) / 3
		area := (math.Pi * raio * math.Sqrt(math.Pow(raio, 2)+math.Pow(altura, 2)))
		fmt.Printf("O volume do cone reto é %.2f\n e a área é %.2f\n ", volume, area)

	case 2:
		volume := (math.Pi * math.Pow(raio, 2) * altura)
		area := 2 * math.Pi * raio * altura
		fmt.Printf("O volume do cilindro é %.2f\n e a área é %.2f\n ", volume, area)

	case 3:
		volume := (4 * math.Pi * math.Pow(raio, 3)) / 3
		area := 4 * math.Pi * math.Pow(raio, 2)
		fmt.Printf("O volume da esfera é %.2f\n e a área é %.2f\n ", volume, area)
	default:
		fmt.Println("Erro, valor icompatível")
	}

}
