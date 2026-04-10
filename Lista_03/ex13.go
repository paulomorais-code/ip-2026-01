package main

import f "fmt"

func main() {
	n := 1.0
	d := 1.0
	var h float64

	for i := 0; i < 50; i++ {
		h += n / d

		n += 2
		d++
	}
	f.Printf("O valor de H é %.4f\n", h)
}
