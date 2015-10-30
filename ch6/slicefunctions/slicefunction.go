package main

import "fmt"

func main() {
	//slice functions append & copy

	//append
	x := []int{1,2,3}
	fmt.Println(x)
	x = append(x, 4, 5, 6, 7)

	fmt.Println(x)

	//copy
	y := make([]int, 4)
	//copy(y[2:], x)
	copy(y, x)
	fmt.Println(y)

	//we can append to x from z
	z := make([]int, 4)
	z[0] = 8
	x = append(x, z...)
	fmt.Println(x)

}
