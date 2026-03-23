package main

import "fmt"

func main() {
	var numero, valor, contador, limite, media float64
	valor = 0
	contador = 1
	limite = 10

	for contador <= limite {
		fmt.Println("Digite o", contador, "° numero:")
		fmt.Scan(&numero)
		valor = valor + numero
		contador = contador + 1
	}
	media = valor / 10
	fmt.Print(media)
}
