package main

import "fmt"

func main() {
	var x [6]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	x[5] = 93

	//with a range
	var t float64 = 0
	for _, v := range x {
		t += v;
	}
	fmt.Println("Range total: ", t / float64(len(x)))
	var total float64 = 0
	for i :=0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))
}
