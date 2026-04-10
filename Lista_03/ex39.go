package main

import f "fmt"

type Boi struct {
	numero int
	peso   float64
}

func main() {
	var bois []Boi
	var b Boi

	for i := 0; i < 90; i++ {

		f.Printf("Boi %d\n", i+1)

		f.Print("Número: ")
		f.Scan(&b.numero)

		f.Print("Peso: ")
		f.Scan(&b.peso)

		bois = append(bois, b)
	}

	maisGordo := bois[0]
	maisMagro := bois[0]

	for _, b := range bois {
		if b.peso > maisGordo.peso {
			maisGordo = b
		}
		if b.peso < maisMagro.peso {
			maisMagro = b
		}
	}

	f.Printf("Boi mais gordo, numero: %d, peso : %f\n", maisGordo.numero, maisGordo.peso)

	f.Printf("Boi mais magro, numero: %d, peso : %f\n", maisMagro.numero, maisMagro.peso)
}
