package main

import f "fmt"

func fatorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * fatorial(n-1)
}

func main() {
	numero := 100.0
	soma := 0.0

	for i := 0; i <= 20; i++ {
		soma += numero / float64(fatorial(i))
		numero--
	}
	f.Println(soma)
}
