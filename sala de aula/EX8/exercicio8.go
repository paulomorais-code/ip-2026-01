package main

import "fmt"

func media(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	soma := 0
	for _, n := range nums {
		soma += n
	}

	return float64(soma) / float64(len(nums))

}
func main() {
	valores := []int(fmt.Scan())
	fmt.Println("Media", media(valores))
}
