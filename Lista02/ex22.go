package main

import "fmt"

func main() {
	var funcionario int
	var horas float64
	fmt.Println("Digite sua matricula:")
	fmt.Scan(&funcionario)

	fmt.Println("Digite suas horas extras")
	fmt.Scan(&horas)
	salarioHE := horas * 10
	salarioB := (3 * 788) + salarioHE

	if salarioB > 1500 {
		inss := salarioB * 0.12
		salarioL := salarioB - inss
		fmt.Println("funcionário", funcionario)
		fmt.Printf("seu salário bruto é %.2f\n", salarioB)
		fmt.Printf("Seu salário liquído é: %.2f\n", salarioL)
	} else if salarioB > 2000 {
		inss := salarioB * 0.12
		Renda := salarioB * 0.2
		salarioL := salarioB - inss - Renda
		fmt.Println("funcionário", funcionario)
		fmt.Printf("seu salário bruto é %.2f\n", salarioB)
		fmt.Printf("Seu salário liquído é:%.2f\n", salarioL)
	}

}
