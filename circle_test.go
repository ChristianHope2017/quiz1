package main

import(
	"testing"
)

func TestArea1(t *testing.T){
Circle := Circle{radius: 2}
    a := Circle.Area1()
	if a != 12.57{
		t.Errorf("Area has FAILED, Expected 16, got %v", a)
	}
}

func TestPerimeter1(t *testing.T){
	Circle := Circle{radius: 2}
		a := Circle.Perimeter1()
		if a != 12.57{
			t.Errorf("Area has FAILED, Expected 16, got %v", a)
		}
	}