package main

import "fmt"

func main() {
	nums, even := halfIt(3)
	fmt.Println(nums, even)
}
//both of those work and are the same thing
func halfIt(number int) (int, bool) {
	return number / 2, number % 2 == 0
//		if (number % 2) == 0 {
//			return number/2, true
//		} else {
//			return number/2, false
//		}
}
