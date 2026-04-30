package main

import f "fmt"

func main() {
	var a int
	var l1 []int
	var p []int
	var i []int
	somap := 0
	somai := 0

	for i := 0; i < 10; i++ {
		f.Printf("Escreva o %d ° numero\n", i+1)
		f.Scan(&a)
		l1 = append(l1, a)
	}
	for _, l := range l1 {
		if l%2 == 0 {
			p = append(p, l)
			somap += l
		} else {
			i = append(i, l)
			somai += l
		}
	}
	f.Printf("Os números pares são: %d e a soma deles é: %d\n", p, somap)
	f.Printf("Os números impares são: %d e a soma deles é: %d\n", i, somai)
}
