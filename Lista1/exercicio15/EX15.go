package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 2; i <= N; i += 2 {
		fmt.Printf("%d^2 = %d\n", i, i*i)
	}
}
