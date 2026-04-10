package main

import f "fmt"

func main() {
	var n int
	var a1, a2 int

	f.Println("Digite o número de termos (N >= 3): ")
	f.Scan(&n)

	f.Println("Digite o primeiro termo: ")
	f.Scan(&a1)

	f.Println("Digite o segundo termo: ")
	f.Scan(&a2)

	serie := make([]int, n)

	serie[0] = a1
	serie[1] = a2

	for i := 2; i < n; i++ {
		if (i+1)%2 == 0 {
			serie[i] = serie[i-1] - serie[i-2]
		} else {
			serie[i] = serie[i-1] + serie[i-2]
		}
	}

	f.Println("Série de Fetuccine:")
	for _, valor := range serie {
		f.Print(valor, " ")
	}
}
