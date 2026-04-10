package main

import f "fmt"

func main() {
	var a, b float64
	a = 1
	b = 15
	soma := 0.0

	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			soma += a / (b * b)
		} else {
			soma -= a / (b * b)
		}
		b--
		a *= 2
	}
	f.Println(soma)
}
