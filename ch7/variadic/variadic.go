package main

import "fmt"

func main() {
	n := average(98,88,92,93,99)
	fmt.Println(n)
}

func average(ys ...float64) float64{
	total := 0.0
	for _, v := range ys {
		total += v
	}
	return total / float64(len(ys))
}