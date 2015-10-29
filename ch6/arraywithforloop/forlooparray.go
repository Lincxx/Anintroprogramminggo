package main

import "fmt"

func main() {
	var x [5]float64

	var total float64 = 0
	for i :=0; i < 5; i++ {
		total += x[i]
	}

	fmt.Println(total /5)
}
