package main

import "fmt"

const (
	kilometer = 1.60934
	fomatter  = "+----------------+"
)

func main() {
	miles, converted := 0., 0.

	fmt.Println("Please enter miles to be converted")
	fmt.Scan(&miles)

	converted = miles * kilometer
	fmt.Println(fomatter)
	fmt.Printf("| Miles: %7.2f |\n", converted)
	fmt.Println(fomatter)

}
