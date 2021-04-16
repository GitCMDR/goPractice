package main

import (
	"fmt"
	"math"
)

/*
-create a type SQUARE ✅
-create a type CIRCLE ✅
-attach a method to each that calculates AREA and returns it
circle area= π r 2 ✅
square area = L * W ✅
-create a type SHAPE that defines an interface as anything that has the AREA method ✅
-create a func INFO which takes type shape and then prints the area ✅???
-create a value of type square ✅
-create a value of type circle ✅
-use func info to print the area of square
-use func info to print the area of circle
 */

type square struct {
	l float64
	w float64
}
type circle struct {
	r float64
}
type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.l * s.w
}

func (c circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	c := circle{
		r: 324.32,
	}

	s := square{
		l: 11,
		w: 23,
	}

	info(c)
	info(s)
}
