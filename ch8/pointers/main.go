package main

import "fmt"

func swap (x *int, y *int) {
	*x = 2
	*y = 1
}

func main() {
	x := 1
	y := 2

	swap(&x, &y)
	fmt.Println(x, y)
}
