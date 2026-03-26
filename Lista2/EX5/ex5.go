package main

import "fmt"

func main() {
	var a int
	fmt.Println("Escreva um numero:")
	fmt.Scan(&a)

	if a%5 == 0 {
		fmt.Println(a, "é divisivel por 5")
	} else {
		fmt.Println(a, "não é divisivel por 5")
	}
}
