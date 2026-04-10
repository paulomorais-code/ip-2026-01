package main

import f "fmt"

func main() {
	var numero, produto int
	f.Println("Escreva um numero inteiro postivo:")
	f.Scan(&numero)
	if numero <= 0 {
		return
	}

	for i := 1; produto < numero; i++ {
		produto = i * (i + 1) * (i + 2)

		if produto == numero {
			f.Printf("O numero %d é triangular\n", numero)
			return
		}
	}
	f.Printf("O numero %d não é triangular\n", numero)
}
