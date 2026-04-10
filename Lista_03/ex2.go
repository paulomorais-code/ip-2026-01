package main

import f "fmt"

func main() {

	soma := 0
	quantidade := 0

	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			soma += i
			quantidade += 1
		}
	}
	media := soma / quantidade

	f.Printf("a soma dos numeros é: %d e a média: %d", soma, media)
}
