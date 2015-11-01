package main

import "fmt"
//this example is not done very often
func main() {
	xs := []float64{98, 93, 77, 82, 83}
	n := average(xs)
	fmt.Println(n)
}

func average(xs []float64) (n float64) {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	n = total / float64(len(xs))
	return n
}