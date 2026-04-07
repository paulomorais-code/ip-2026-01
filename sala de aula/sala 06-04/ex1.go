package main

import f "fmt"

func peso(alt float64) float64 {
	return (72.7 * alt) - 58
}

type pessoa struct {
	nome   string
	altura float64
	peso   float64
}

func main() {
	var altura float64
	var nome string
	var pes []pessoa

	for {
		f.Println("Escreva seu nome ou FIM para encerrar")
		f.Scan(&nome)
		if nome == "FIM" {
			break
		}

		f.Println("Qual a sua altura")
		f.Scan(&altura)

		Novapessoa := pessoa{
			nome:   nome,
			altura: altura,
			peso:   peso(altura),
		}

		pes = append(pes, Novapessoa)
	}

	for _, p := range pes {
		f.Printf("Nome: %s | Peso ideal, %2.2f\n", p.nome, p.peso)
	}
}
