package main

import "fmt"


func main() {
	//we don't pass the fun anything
	x := 0
	increment := func() int {
		//we are referencing x here, this is because of the scope
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
