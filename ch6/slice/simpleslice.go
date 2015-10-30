package main

import "fmt"

func main() {
	//what is the difference. There isn't a length with a slice [] <- is empty
	var x = []float64{1,2,3}
	//create a new slice
	var y = x[1:]
	fmt.Println(x,y)

	//creates a slice, here we set the length of the slice
	var t = make([]float64, 10)
	fmt.Println(t)
	//creates a slice, here we set the length to 15 and makes the segment go to 10, so we have 5 extra at the end
	//This is not used very often
	var r = make([]float64, 10, 15)
	fmt.Println(r)

	//convert an array to a slice
	var arr = [4]float64{1,2,3,4}
	var sliarr = arr[0:2]
	fmt.Println(sliarr)
}
