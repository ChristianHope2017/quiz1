package main

import(
	"math"
)

type Triangle struct{
	base float64
	height float64
}

func (t Triangle) Area() float64{
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64{
	return math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
}


func main(){

	triangle := Triangle{
		base: 4,
		height: 3,
	}

	triangle.Area()
	triangle.Perimeter()
}