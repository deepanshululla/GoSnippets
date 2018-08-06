package main

import "fmt"

type shape interface {
	printArea() float64
}
type square struct {
	sideLength float64;
}

type triangle struct {
	base float64
	height float64
}
func (s square) printArea() float64{
	area:=s.sideLength*s.sideLength;
	fmt.Println(area)
	return area;
}

func (t triangle) printArea() float64{
	area:=0.5*t.base*t.height;
	fmt.Println(area);
	return area;
}

func main() {
	t1:=triangle{base:5,height:6};
	t1.printArea();
	s:=square{sideLength:5}
	s.printArea();
}
