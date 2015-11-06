package main

import "fmt"

//this star *int is the pointer, that points back to the memory location
func zero(xPtr *int) {
	fmt.Println(xPtr)
	*xPtr = 0
}
func main() {
	x := 5
	//this & is the memory address
	zero(&x)
	fmt.Println(&x)
	fmt.Println(x) // x is 0
}
