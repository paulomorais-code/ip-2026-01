package main

import f "fmt"

func main() {
	soma := 0
	var a int
	var lista1 []int
	var lista2 []int
	var lp []int
	var li []int
	for i := 0; i < 10; i++ {
		f.Println("Digite um numero para l1:")
		f.Scan(&a)
		lista1 = append(lista1, a)
	}

	for j := 0; j < 5; j++ {
		f.Println("Digite um numero para l2:")
		f.Scan(&a)
		lista2 = append(lista2, a)
	}

	for _, l2 := range lista2 {
		soma += l2
	}

	for _, l1 := range lista1 {
		if l1%2 == 0 {
			num := l1 + soma
			lp = append(lp, num)
		} else {
			num := l1 + soma
			li = append(li, num)
		}
	}
	f.Printf("Os vetores resultantes são: %d, %d\n", lp, li)
}
