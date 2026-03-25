package main

import "fmt"

func main() {
	var nota1, p1, nota2, p2 int
	fmt.Println("Escreva a primeira nota e seu peso")
	fmt.Scan(&nota1, &p1)
	fmt.Println("Escreva a segunda nota e seu peso")
	fmt.Scan(&nota2, &p2)

	media := ((nota1 * p1) + (nota2 * p2)) / (p1 + p2)
	fmt.Println(media)
}
