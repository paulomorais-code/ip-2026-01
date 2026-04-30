package main

import f "fmt"

type LinhaTabela struct {
	Nota int
	FA   int
	FR   float64
}

func main() {
	var notas [15]int
	contagem := make(map[int]int)

	for i := 0; i < 15; i++ {
		f.Printf("Escreva o %d ° numero\n", i+1)
		f.Scan(&notas[i])
		if notas[i] > 10 || notas[i] < 0 {
			f.Println("nota inválida")
			return
		}

		contagem[notas[i]]++
	}

	var tabela []LinhaTabela

	for i := 0; i <= 10; i++ {
		fa := contagem[i]
		fr := float64(fa) / 15.0

		tabela = append(tabela, LinhaTabela{Nota: i, FA: fa, FR: fr})
	}

	f.Printf("%s | %s | %s\n", "Nota", "FA", "FR")
	for _, linha := range tabela {
		f.Printf("%d | %d | %.2f\n", linha.Nota, linha.FA, linha.FR)
	}
}
