package main

import "fmt"
import "math"

type Shape interface {
	area() float64
}

type Rectangle struct {
	x1, x2, y1, y2  float64
}


type Circle struct  {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64{
	return math.Pi * c.r * c.r
}


func(r *Rectangle) area() float64{
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)

	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func main() {

//	var s Shape
//	s = c
//	s = r

	r := Rectangle{x1: 2, x2: 3, y1: 2, y2: 3}
	fmt.Println()
	fmt.Println(r.area())

	//var c Circle

	c := Circle{x: 0, y: 0, r: 5}
	//or like this if we know the order
	//c := Circle{0,0,5}

	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c.area())
	c.x = 10
	c.y = 5

}

