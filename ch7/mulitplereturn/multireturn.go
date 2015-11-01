package main

import "fmt"
//this doesn't work????
func main() {
	x, y := f(1, 2)
	fmt.Println(x, y)
}

func f(int, int){
	return 5, 6
}
