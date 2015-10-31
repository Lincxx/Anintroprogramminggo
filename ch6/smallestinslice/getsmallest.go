package main

import "fmt"
import "sort"

//find the smallest
func main() {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	//first solution
//	min := x[0]
//
//	for _, v := range x {
//		if v < min{
//			min = v
//		}
//	}

	sort.Ints(x) //first solution would be faster
	fmt.Println(x[0])
}
