package main

import "fmt"

func main() {
	var quantidade int
	var soma, nota float64
	soma = 0
	fmt.Println("Quantas notas serão coletadas?")
	fmt.Scan(&quantidade)

	for i := 0; i < quantidade; i++ {
		fmt.Println("Nota:")
		fmt.Scan(&nota)
		soma += nota
	}
	media := soma / float64(quantidade)

	fmt.Printf("média é = %.2f\n", media)

}
