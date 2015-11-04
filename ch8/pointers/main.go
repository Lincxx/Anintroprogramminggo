package main

import "fmt"

func swap (x, y *int) {
	//we store a *x in a tmp var
	var tmp = *x
	*x = *y
	*y = tmp

	//short form *x, *y = *y, *x
}

func main() {
	x := 1
	y := 2

	swap(&x, &y)
	fmt.Println(x, y)
}
