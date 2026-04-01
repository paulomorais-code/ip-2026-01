package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Escreva os coeficientes A, B e C nessa ordem:")
	fmt.Scan(&a, &b, &c)

	if a == 0 {
		fmt.Println("não é função de 2° grau")
	}

	delta := (b * b) - (4 * a * c)

	switch {
	case delta < 0:
		fmt.Println("Classificação: RAÍZES IMAGINÁRIAS")
		fmt.Println("Não existem raízes reais para esta equação.")

	case delta == 0:
		x := -b / (2 * a)
		fmt.Println("Classificação: RAIZ ÚNICA")
		fmt.Printf("a raiz é %.2f\n", x)

	case delta > 0:
		sqrtdelta := math.Sqrt(delta)
		x1 := (-b + sqrtdelta) / (2 * a)
		x2 := (-b - sqrtdelta) / (2 * a)
		fmt.Println("classificação: possui duas raizes")
		fmt.Printf("as raiz 1 é %.2f\n", x1)
		fmt.Printf("as raiz 2 é %.2f\n", x2)
	}
}
