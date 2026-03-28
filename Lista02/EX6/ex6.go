package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Escreva dois números:")
	fmt.Scan(&a, &b)

	if a%b == 0 {
		fmt.Println(a, "é divisivel por", b)
	} else {
		fmt.Println(a, "não é divisivel por", b)
	}
}
