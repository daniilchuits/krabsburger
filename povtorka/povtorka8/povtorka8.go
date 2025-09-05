package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	Base, Height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	fig := []Shape{Circle{5}, Rectangle{5, 5}, Triangle{5, 5}}
	for _, f := range fig {
		PrintArea(f)
	}
}
