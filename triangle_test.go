package main

import(
	"testing"
)

func TestArea(t *testing.T){
triangle := Triangle{base: 4, height: 3}
    a := triangle.Area()
	if a != 6{
		t.Errorf("Area has FAILED, Expected 6, got %v", a)
	}
}

func TestPerimeter(t *testing.T){
	triangle := Triangle{base: 4, height: 3}
    a := triangle.Perimeter()
		if a != 12{
			t.Errorf("Area has FAILED, Expected 12, got %v", a)
		}
	}