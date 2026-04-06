package main

import f "fmt"

func peso(alt float64) float64 {
	return (72.7 * alt) - 58
}

func main() {
	f.Println("Fim quando digitar FIM")
	type pessoa struct {
		nome   string
		altura float64
		peso   float64
	}
	var pes = pessoa{}

	for pes.nome != "FIM" {
		f.Println("Escreva seu nome")
		f.Scan(&pes.nome)
		if pes.nome == "FIM" {
			return
		}

		f.Println("Qual a sua altura")
		f.Scan(&pes.altura)

		f.Printf("O peso ideal de %s é: %.2f\n", pes.nome, peso(pes.altura))
	}
	println("FIM")

}
