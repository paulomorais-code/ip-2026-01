package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Println("Unidade; dezena; centena")
	fmt.Scan(&n1, &n2, &n3)
	if n1 < 0 || n1 > 9 || n2 < 0 || n2 > 9 || n3 < 0 || n3 > 9 {
		fmt.Println("Digito inválido")
		return
	}
	numero := n1*100 + n2*10 + n3
	quadrado := numero * numero
	fmt.Println(numero, quadrado)
}
