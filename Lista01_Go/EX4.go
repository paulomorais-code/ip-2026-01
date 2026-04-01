package main

import "fmt"

func main() {
	var salariomin, Consumo float64
	fmt.Println("salariomin, consumo")
	fmt.Scan(&salariomin, &Consumo)

	Kw := ((70 * salariomin) / 100) / 100

	CustoConsumo := Kw * Consumo

	Custodesconto := CustoConsumo * 0.9

	fmt.Printf("Custo por Kw: R$: %.2f\n", Kw)
	fmt.Printf("Custo do consumo: R$:%.2f\n", CustoConsumo)
	fmt.Printf("Custo com desconto:R$:%.2f\n", Custodesconto)

}
