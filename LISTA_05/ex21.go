package main

import f "fmt"

func inverter(l []int) {
	n := len(l) - 1
	if n <= 0 {
		return
	}
	l[0], l[n] = l[n], l[0]
	inverter(l[1:n])
}

func main() {
	var a, b int
	var lista []int

	for i := 0; i < 10; i++ {
		f.Println("Escreva um valor para lista1:")
		f.Scan(&b)
		lista = append(lista, b)
	}

	f.Println("Escolha uma opção 0-sair 1-direta 2-inverter")
	f.Scan(&a)
	if a < 0 || a > 2 {
		return
	}

	switch a {
	case 0:
		{
			return
		}

	case 1:
		{
			f.Println(lista)
		}

	case 2:
		{
			inverter(lista)
			f.Println(lista)
		}
	}
}
