package main

import f "fmt"

func main() {
	var conta [10]int
	var saldo [10]float64
	var sel, comparador int
	var saque float64

	for i := 0; i < 10; i++ {
		f.Printf("Escreva a %d° conta:", i+1)
		f.Scan(&conta[i])
		j := i - 1

		for j >= 0 {
			if conta[j] == conta[i] {
				f.Println("Erro")
				return
			}
			j--
		}
		f.Printf("Escreva o %d° saldo:", i+1)
		f.Scan(&saldo[i])
	}
	for {
		f.Println("Escolha um opcão")
		f.Println(" (1- Para efetuar depósito 2- Efetuar saque 3-Consultar ativo bancário 4- Finalizar)")
		f.Scan(&sel)

		if sel == 4 {
			break
		}

		switch sel {

		// caso depósito
		case 1:
			var status bool
			var valor int
			var dep float64
			f.Println("Escreva a conta")
			f.Scan(&comparador)

			for i, c := range conta {
				if c == comparador {
					status = true
					valor = i
					break
				}
			}
			if status {
				f.Println("Conta encontrada!")
				f.Println("Qual o valor a ser depositado? ")
				f.Scan(&dep)
				saldo[valor] = saldo[valor] + dep
				f.Printf("Novo saldo: %.2f\n", saldo[valor])
			} else {
				f.Println("Conta não encontrada!")
			}

			// caso saque

		case 2:
			var status bool
			var valor int
			f.Println("Escreva a conta")
			f.Scan(&comparador)

			for i, c := range conta {
				if c == comparador {
					status = true
					valor = i
					break
				}
			}

			if status {
				f.Println("Conta encontrada!")
				f.Printf("Seu saldo é %.2f\n", saldo[valor])
				f.Print("Qual o valor a ser sacado?")
				f.Scan(&saque)

				//verificar saldo

				if saque <= saldo[valor] {
					saldo[valor] = saldo[valor] - saque
					f.Printf("Saque de %.2f realizado com sucesso\n", saque)
					f.Printf("Novo saldo: %.2f\n", saldo[valor])
				} else {
					f.Println("Saldo insuficiente")
				}
			} else {
				f.Println("Conta não encontrada!")
			}

			// caso ativo bancário

		case 3:
			soma := 0.0

			for _, s := range saldo {
				soma += s
			}
			f.Printf("O total de ativo bancário é %.2f\n", soma)

		}
	}
}
