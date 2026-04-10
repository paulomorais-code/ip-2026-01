package main

import (
	"fmt"
)

func main() {
	var cpf string

	fmt.Print("Digite o CPF (somente números): ")
	fmt.Scan(&cpf)

	if len(cpf) != 11 {
		fmt.Println("CPF inválido")
		return
	}

	// Verifica se todos os dígitos são iguais
	igual := true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			igual = false
			break
		}
	}
	if igual {
		fmt.Println("CPF inválido")
		return
	}

	// 1º dígito
	soma := 0
	for i := 0; i < 9; i++ {
		soma += int(cpf[i]-'0') * (10 - i)
	}

	resto := (soma * 10) % 11
	if resto == 10 {
		resto = 0
	}

	if resto != int(cpf[9]-'0') {
		fmt.Println("CPF inválido")
		return
	}

	// 2º dígito
	soma = 0
	for i := 0; i < 10; i++ {
		soma += int(cpf[i]-'0') * (11 - i)
	}

	resto = (soma * 10) % 11
	if resto == 10 {
		resto = 0
	}

	if resto != int(cpf[10]-'0') {
		fmt.Println("CPF inválido")
		return
	}

	fmt.Println("CPF válido")
}
