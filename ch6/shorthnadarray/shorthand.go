package main

import "fmt"

func main() {
	//this is the short
	var arr = [6]float64{
		98,
		76,
		45,
		89,
		99,
		100,
	}

	var total = 0.0

	for _, v := range arr {
		total += v
	}

	fmt.Println( total / float64(len(arr)))
}
