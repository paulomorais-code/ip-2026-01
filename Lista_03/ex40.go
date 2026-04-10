package main

import (
	f "fmt"
)

func main() {
	preco := 6.0
	quantidade := 130.0
	custoFixo := 300.0
	variacaoPreco := 0.60
	variacaoQtd := 30.0

	var maxLucro, maxPreco, maxQtd float64

	f.Printf("%s | %s | %s\n", "Preço (R$)", "Qtd Ingressos", "Lucro (R$)")

	for p := preco; p >= 1.0; p -= variacaoPreco {
		receita := p * quantidade
		lucro := receita - custoFixo

		f.Printf("R$ %.2f | %.2f | R$ %.2f\n", p, quantidade, lucro)

		if lucro > maxLucro {
			maxLucro = lucro
			maxPreco = p
			maxQtd = quantidade
		}

		quantidade += variacaoQtd
	}

	f.Printf("\nLUCRO MÁXIMO ESPERADO: R$ %.2f\n", maxLucro)
	f.Printf("PREÇO DO INGRESSO: R$ %.2f\n", maxPreco)
	f.Printf("NÚMERO DE INGRESSOS: %.0f\n", maxQtd)
}
