package main

import f "fmt"

func soma(l []float64) float64 {
	n := len(l) - 1
	if n == 0 {
		return l[n]
	}
	return l[n] + soma(l[:n])

}

func main() {
	reais := []float64{5.5, 5.5, 6.0}
	f.Println("A soma é", soma(reais))
}
