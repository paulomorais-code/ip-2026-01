package main

import (
	f "fmt"
	m "math"
)

func main() {
	raio := 0.5

	for i := 0; i < 40; i++ {
		volume := (4 * m.Pi * raio * raio * raio) / 3
		f.Printf("raio: %.2f | volume:%.2f\n", raio, volume)
		raio += 0.5
	}
}
