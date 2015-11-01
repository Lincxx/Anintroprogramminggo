package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	//10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(10))
}
