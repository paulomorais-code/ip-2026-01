package main

import f "fmt"

func main() {
	var b, n int
	f.Println("Informe a base e o expoente:")
	f.Scan(&b, &n)
	potencia := 1

	for i := 1; i <= n; i++ {
		potencia *= b
	}
	f.Printf("%d ^ %d é igual a: %d", b, n, potencia)
}
