package main

import(
	"math"
)

type Circle struct{
	radius float64
}

func (c Circle) Area1() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter1() float64{
	return 2 * math.Pi * c.radius
}
