package main

import "fmt"

func main() {
	var numero int

	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Println("O NUMERO É DIVISIVEL")
	} else {
		fmt.Println(" O NUMERO NÃO E DIVISIVEL")
	}
}
