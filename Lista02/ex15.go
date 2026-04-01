package main

import "fmt"

func main() {
	var a, m, d int
	var mes string
	fmt.Println("Escreva o ano:")
	fmt.Scan(&a)

	fmt.Println("Escreva o mês:")
	fmt.Scan(&m)

	fmt.Println("Escreva o dia:")
	fmt.Scan(&d)

	if m == 1 {
		mes = "janeiro"
	} else if m == 2 {
		mes = "fevereiro"
	} else if m == 3 {
		mes = "março"
	} else if m == 4 {
		mes = "abril"
	} else if m == 5 {
		mes = "maio"
	} else if m == 6 {
		mes = "junho"
	} else if m == 7 {
		mes = "julho"
	} else if m == 8 {
		mes = "agosto"
	} else if m == 9 {
		mes = "setembro"
	} else if m == 10 {
		mes = "outubro"
	} else if m == 11 {
		mes = "novembro"
	} else if m == 12 {
		mes = "dezembro"
	} else {
		mes = "não existe"
	}
	fmt.Println(d, "de", mes, "de", a)
}
