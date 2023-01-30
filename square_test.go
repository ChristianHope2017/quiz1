package main

import(
	"testing"
)

func TestSquare(t *testing.T){
	a, b := Square(4)
	if a != 16{
		t.Errorf("Area has FAILED, Expected 16, got %v", a)
	}

	if b != 16{
		t.Errorf("Perimeter has FAILED, Expected 16, got %v", b)
	}
}