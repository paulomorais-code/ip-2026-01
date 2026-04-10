package main

import f "fmt"

type pessoa struct {
	idade int

	altura float64

	peso float64
}

func main() {
	var a, idade int
	var altura, peso, soma, media, porcentagem, Cont float64
	var pes []pessoa
	ContI := 0
	var ContC float64 = 0
	ContP := 0
	for {
		f.Println("você quer digitar mais dados? 1 - Sim, 2- Não")
		f.Scan(&a)

		if a == 2 {
			break
		}

		f.Println("Qual sua idade, sua altura e seu peso, respectivamente?")
		f.Scan(&idade, &altura, &peso)

		Novapessoa := pessoa{
			idade:  idade,
			altura: altura,
			peso:   peso,
		}

		pes = append(pes, Novapessoa)
	}

	for _, p := range pes {
		Cont++
		if p.idade > 50 {
			ContI++
		}
		if p.idade >= 10 && p.idade <= 20 {
			soma = soma + p.altura
			ContC++
			media = soma / ContC
		}
		if p.peso <= 40 {
			ContP++
		}
		porcentagem = float64(ContP) / Cont
	}

	f.Printf("a quantidade de pessoas mais velhas que 50 anos é %d\n", ContI)
	f.Printf("A média de altura de pessoas entre 10 e 20 anos é %.2f\n", media)
	f.Printf("A porcentagem de pessoas com menos de 40 kg é %.2f\n", porcentagem)
}
