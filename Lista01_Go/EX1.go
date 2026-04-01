package main

import "fmt"

func main() {
	var nota1, nota2, nota3, media float64
	fmt.Print("digite a primeira nota:")
	fmt.Scanln(&nota1)

	fmt.Print("digite a segunda nota:")
	fmt.Scanln(&nota2)

	fmt.Print("digite a terceira nota:")
	fmt.Scanln(&nota3)

	media = (nota1 + nota2 + nota3) / 3
	if media >= 6 {
		fmt.Println(media)
		fmt.Println("APROVADO")
	} else {
		fmt.Println(media)
		fmt.Println("REPROVADO")
	}
}
