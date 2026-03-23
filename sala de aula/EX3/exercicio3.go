package main

import "fmt"

func main() {
	var preco, desc float64
	fmt.Scan(&preco, &desc)
	novopreco := preco * desc
	fmt.Println(novopreco)
}
