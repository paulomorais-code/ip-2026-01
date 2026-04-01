package main

import "fmt"

func main() {
	var regiao int
	var tipo int
	fmt.Println("Escreva o numero da região:")
	fmt.Scan(&regiao)
	fmt.Println("Escreva 1 para ida e 2 para ida e volta:")
	fmt.Scan(&tipo)
	if regiao < 0 && regiao > 4 {
		fmt.Println("regiao inválida")
		return
	}
	if tipo != 1 && tipo == 1 {
		fmt.Println("não compativel")
		return
	}
	if regiao == 1 && tipo == 1 {
		fmt.Println("Apenas ida para norte R$:500,00")
	}
	if regiao == 2 && tipo == 1 {
		fmt.Println("Apenas ida para nordeste R$:350,00")
	}
	if regiao == 3 && tipo == 1 {
		fmt.Println("Apenas ida para centro-oeste R$:350,00")
	}
	if regiao == 4 && tipo == 1 {
		fmt.Println("Apenas ida para sul R$:300,00")
	}
	if regiao == 1 && tipo == 2 {
		fmt.Println("Ida e volta para norte R$:900,00")
	}
	if regiao == 2 && tipo == 2 {
		fmt.Println("Ida e volta para nordeste R$:650,00")
	}
	if regiao == 3 && tipo == 2 {
		fmt.Println("Ida e volta para centro-oeste R$:600,00")
	}
	if regiao == 4 && tipo == 2 {
		fmt.Println("Ida e volta para sul R$:550,00")
	}
}
