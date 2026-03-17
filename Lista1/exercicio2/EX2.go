package main

import "fmt"

func main() {
	var casos int
	fmt.Scan(&casos)
	fmt.Print("Numero do jogo:")
	for i := 1; i <= casos; i++ {
		var publico int
		var pPop, pGeral, pArq, pCad float64
		fmt.Println("publico total, publico Popular, publico Geral, publico Arquibancada, publico Cadeira:")
		fmt.Scan(&publico, &pPop, &pGeral, &pArq, &pCad)
		popular := float64(publico) * (pPop / 100)
		Geral := float64(publico) * (pGeral / 100)
		Arq := float64(publico) * (pArq / 100)
		Cad := float64(publico) * (pArq / 100)

		Renda := popular*1 + Geral*5 + Arq*10 + Cad*20
		fmt.Printf(" A RENDA DO JOGO N. %d = %.2f\n", i, Renda)

	}
}
