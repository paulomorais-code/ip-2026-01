package main

import f "fmt"

func main() {
	var numero int

	for {
		i := 1
		f.Println("Digite um numero ( e um numero <= 0 para sair):")
		f.Scan(&numero)

		if numero <= 0 {
			return
		}

		for (i*i) != numero && i < numero {
			i++
		}
		if i*i == numero {
			f.Printf("%d é quadrado perfeito\n", numero)

		} else {
			f.Printf(" %d não é quadrado perfeito\n", numero)

		}

	}
}
