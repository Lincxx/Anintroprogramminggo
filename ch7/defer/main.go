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
	//defer means to make something go last
	defer fmt.Println(factorial(10))
	fmt.Println("Hello")

}
