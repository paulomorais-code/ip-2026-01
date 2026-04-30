package main

import f "fmt"

func main() {
	var a int
	var l1 []int
	var l2 []int
	var lr []int

	for i := 0; i < 10; i++ {
		f.Printf("Escreva o %d ° numero\n", i+1)
		f.Scan(&a)
		l1 = append(l1, a)
	}

	for i := 0; i < 10; i++ {
		f.Printf("Escreva o %d ° numero\n", i+1)
		f.Scan(&a)
		l2 = append(l2, a)
	}

	for i := 0; i < 10; i++ {
		lr = append(lr, l1[i], l2[i])
	}
	f.Println(lr)
}
