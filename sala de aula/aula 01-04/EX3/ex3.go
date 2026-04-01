package main

import "fmt"

func media(x, y, z float64) float64 {
	var media float64

	media = (x + y + z) / 3
	return media

}
func main() {
	var x, y, z float64
	fmt.Println("Escreva 3 números: ")
	fmt.Scan(&x, &y, &z)

	fmt.Printf("%.2f\n", media(x, y, z))
}
