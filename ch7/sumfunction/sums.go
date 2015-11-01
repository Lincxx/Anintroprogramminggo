package main

import "fmt"

func main() {
	//I'm using a slice of number
	nums := []int{1,2,3,4,5}
	mySumSlice := sumSlice(nums)
	fmt.Println(mySumSlice)

	mySum := sum(1,2,3,4,5)
	fmt.Println(mySum)


}

func sumSlice(numbers []int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

func sum(numbers ...int) int{
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}

