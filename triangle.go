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
	a := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	return a + t.base + t.height
}


func main(){

	triangle := Triangle{
		base: 4,
		height: 3,
	}

	circle := Circle{
		radius: 4,
	}

	triangle.Area()
	triangle.Perimeter()

	circle.Area1()
	circle.Perimeter1()

	Square(4)
}