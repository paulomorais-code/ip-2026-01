package main

import "fmt"

func main() {
	var valori, a, b, c, d float64
	fmt.Println("Escreva o valor do carro:")
	fmt.Scan(&valori)
	fmt.Print("deseja ter Ar condicionado? digite 1 para sim e 0 para não: ")
	fmt.Scan(&a)
	fmt.Print("deseja ter Pintura metálica? digite 1 para sim e 0 para não: ")
	fmt.Scan(&b)
	fmt.Print("deseja ter Vidro elétrico? digite 1 para sim e 0 para não: ")
	fmt.Scan(&c)
	fmt.Print("deseja ter direção hidráulica? digite 1 para sim e 0 para não: ")
	fmt.Scan(&d)

	if a == 1 {
		valori = valori + 1750
	}

	if b == 1 {
		valori = valori + 800
	}

	if c == 1 {
		valori = valori + 1200
	}

	if d == 1 {
		valori = valori + 2000
	}

	fmt.Println("o valor total é", valori)
}
