package main

import "fmt"
//a function that returns a function
func increasingNumbersGenerator() func() int {
	x := 0
	return func () int {
	//we are referencing x here, this is because of the scope
	x++
	return x
	}
}

func main() {
	//the increasingNumbersGenerator function has access to the x varible
	numbers := increasingNumbersGenerator
	fmt.Println(numbers())
	fmt.Println(numbers())
}
