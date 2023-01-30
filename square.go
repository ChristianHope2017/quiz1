package main

import(
	"math"
)


func Square(x float64) (float64, float64){
	Area := math.Pow(x, 2)
	Perimeter := 4 * x
	return Area, Perimeter
}
