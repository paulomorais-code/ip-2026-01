package main

import f "fmt"

func main() {
	soma := 0
	j := 1
	for i := 37; i > 0; i-- {
		soma += (i * (i + 1)) / j
		j++
	}
	f.Print(soma)
}
