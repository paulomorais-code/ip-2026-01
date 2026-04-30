package main

import f "fmt"

func ordenar(arr []pessoa) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j].meses > arr[j+1].meses {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

type pessoa struct {
	numero int
	meses  int
}

func main() {
	var pessoas []pessoa
	var a int
	var b pessoa

	f.Println("Quantas pessoas serão cadastradas? (max 100)")
	f.Scan(&a)

	if a > 100 {
		f.Println("número inválido")
		return
	}

	for i := 0; i < a; i++ {
		f.Println("Escreva o numero do empregado e seu tempo em meses")
		f.Scan(&b.numero, &b.meses)
		pessoas = append(pessoas, b)
	}

	ordenar(pessoas)
	f.Println("Os funcionários mais novos são:")
	f.Println(pessoas[0].numero, pessoas[1].numero, pessoas[2].numero)
}
