package main

import f "fmt"

func main() {
	var numero int
	var numeros []int

	// Leitura
	for {
		f.Println("Digite números (30000 para parar):")
		f.Scan(&numero)

		if numero == 30000 {
			break
		}

		numeros = append(numeros, numero)
	}

	soma := 0
	maior := numeros[0]
	menor := numeros[0]

	somaPares := 0
	qtdPares := 0
	qtdImpares := 0

	for _, n := range numeros {
		soma += n

		if n > maior {
			maior = n
		}
		if n < menor {
			menor = n
		}

		if n%2 == 0 {
			somaPares += n
			qtdPares++
		} else {
			qtdImpares++
		}
	}

	quantidade := len(numeros)
	media := float64(soma) / float64(quantidade)

	var mediaPares float64
	if qtdPares > 0 {
		mediaPares = float64(somaPares) / float64(qtdPares)
	}

	percentualImpares := float64(qtdImpares) / float64(quantidade) * 100

	f.Println("Soma:", soma)
	f.Println("Quantidade:", quantidade)
	f.Println("Média:", media)
	f.Println("Maior número:", maior)
	f.Println("Menor número:", menor)
	f.Println("Média dos pares:", mediaPares)
	f.Printf("Percentual de ímpares: %.2f%%\n", percentualImpares)
}
