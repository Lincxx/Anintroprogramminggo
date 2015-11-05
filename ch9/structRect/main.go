package main

import (
"fmt"
"math"
)

type Rectangle struct {
	x1, x2, y1, y2  float64
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


	r := Rectangle{x1: 2, x2: 3, y1: 2, y2: 3}
	fmt.Println()
	fmt.Println(r.area())

}
