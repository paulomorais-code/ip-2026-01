package main

import f "fmt"

func main() {
	var n1, n2 int

	f.Println("Escreva dois numeros:")
	f.Scan(&n1, &n2)
	multiplicação := 0

	for i := 0; i < n2; i++ {
		multiplicação += n1
	}
	f.Printf("%d x %d = %d", n1, n2, multiplicação)
}
