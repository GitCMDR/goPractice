package main

import "fmt"

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
pass a func into a func as an argument
*/

func main() {

	v3 := multiplyNot5(multiply, []int{1,2,3,4,5,6}...)
	fmt.Println(v3)

}

func multiply(x3 ...int) int {
	var figure int
	for _, v := range x3 {
		figure += v
	}
	return figure
}

func multiplyNot5(f func(x3 ...int) int, v3 ...int) int {
	var yi []int

	for _, v := range v3 {
		if v != 5 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
