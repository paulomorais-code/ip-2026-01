package main

import "fmt"

func binario(n int) string {
	if n == 0 {
		return ""
	}

	if n%2 == 0 {
		return binario(n/2) + "0"
	} else {
		return binario(n/2) + "1"
	}
}

func main() {
	var n int
	fmt.Println("Escreva um numero:")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("0")
	} else {
		fmt.Println(binario(n))
	}
}
