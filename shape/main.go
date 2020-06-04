package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println("Area of", reflect.TypeOf(s), ":", s.getArea())
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	result := t.base * t.height * 0.5
	return result
}

type square struct {
	edge float64
}

func (s square) getArea() float64 {
	return s.edge * s.edge
}

func main() {
	t := triangle{base: 2, height: 3}
	printArea(t)

	s := square{edge: 2}
	printArea(s)
}
