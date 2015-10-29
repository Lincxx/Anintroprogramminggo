package main

import "fmt"

func main() {
	fmt.Print("Enter a number to be converted to Celsius")
	f := 0.
	fmt.Scanf("%f", &f)
	println(f)

	c := (f - 32) * 5 / 9

	fmt.Println(c)

}
