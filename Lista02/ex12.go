package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Escreva sua idade:")
	fmt.Scan(&idade)

	if idade < 0 {
		fmt.Println("não existe")
		return
	}
	if idade <= 2 {
		fmt.Println("Recém-nascido")
	} else if idade <= 11 {
		fmt.Println("Criança")
	} else if idade <= 19 {
		fmt.Println("Adolescente")
	} else if idade <= 55 {
		fmt.Println("Adulto")
	} else {
		fmt.Println("Idoso")
	}
}
