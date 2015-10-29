package main

import (
	"fmt"
	"strconv"
)
import "os"

const (
	kilometer = 1.60934
)

func main() {

	miles, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err)
	}

	//miles, _ := 0., 0.
	//fmt.Println("Please enter miles to be converted")
	//fmt.Scanln(&miles)
	//converted = miles * kilometer

	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html>")
	fmt.Println(" <head></head>")
	fmt.Println(" <body>")

	fmt.Printf(" Miles: %7.2f<br> \n", miles)
	fmt.Printf(" Kilomoters: %7.2f<br> \n", miles*kilometer)

	fmt.Println(" </body>")
	fmt.Println("</html>")

}
