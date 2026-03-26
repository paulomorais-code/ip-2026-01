package main

import "fmt"

func main() {
	var av1, av2, av3, atv float64
	fmt.Println("Escreva as notas, 1, 2 e 3 e a média das atividades")
	fmt.Scan(&av1, &av2, &av3, &atv)

	media := (av1 + (av2 * 2) + (av3 * 3) + atv) / 7

	if media < 4 {
		fmt.Printf("conceito E, média %.2f\n", media)
	} else if media < 6 {
		fmt.Printf("conceito D, média %.2f\n", media)
	} else if media < 7.5 {
		fmt.Printf("conceito C, média %.2f\n", media)
	} else if media < 9 {
		fmt.Printf("conceito B, média %.2f\n", media)
	} else if media <= 10 {
		fmt.Printf("conceito A, média %.2f\n", media)
	}
}
