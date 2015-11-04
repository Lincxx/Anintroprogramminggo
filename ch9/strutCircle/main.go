package main

import "fmt"
import "math"

type Circle struct  {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64{
	return math.Pi * c.r * c.r
}

func main() {
	//var c Circle

	c := Circle{x: 0, y: 0, r: 5}
	//or like this if we know the order
	//c := Circle{0,0,5}

	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c.area())
	c.x = 10
	c.y = 5

}
