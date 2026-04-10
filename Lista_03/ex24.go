package main

import (
	f "fmt"
	"math"
)

func main() {

	f.Printf("Ângulo (A) | Seno (Maclaurin)\n")
	f.Println("-------------------------------------------")

	for i := 0; i <= 63; i++ {
		a := float64(i) / 10.0

		a3 := math.Pow(a, 3) / 6.0
		a5 := math.Pow(a, 5) / 120.0
		a7 := math.Pow(a, 7) / 5040.0

		senA := a - a3 + a5 - a7

		f.Printf("%-10.1f | %-20.6f\n", a, senA)
	}
}
