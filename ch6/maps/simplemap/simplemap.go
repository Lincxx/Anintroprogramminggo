package main

import "fmt"

func main() {
	var x  = make(map[string]int)

	x["anything"] = 2

	//maps can return two value, the value itself along with a true or false return
	x["anything"] = 0
	fmt.Println(x["anything"])

	v, ok :=x["something else"]

	fmt.Println(v, ok)

	//shorthand map
	var z = map[int]int{
		7:10,
		3:8,
		2:2,
	}
	fmt.Println(z)
	//remove item from a map with delete

	delete(z, 2)
	fmt.Println(z)


}