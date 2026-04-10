package main

import f "fmt"

func main() {
	soma := 0
	for i := 1; i <= 20; i++ {
		f.Println(i)
		soma = soma + i
	}
	f.Println(soma)
}
