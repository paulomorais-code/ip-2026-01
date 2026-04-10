package main

import f "fmt"

func main() {
	var a, b int

	resultado := 1
	f.Println("Digite a base:")
	f.Scan(&a)

	f.Println("Digite a exponte:")
	f.Scan(&b)

	for i := 1; i <= b; i++ {
		resultado = resultado * a
	}

	f.Printf("O resultado de %d ^ %d é %d\n", a, b, resultado)
}
