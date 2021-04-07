package main

// looks like go arrays are not as widely used btw, but cool content to know, slices is the preferred way to work when wanting to play with array-like structures

import "fmt"

func main() {
	var x [5]int // for arrays in go you need to specify type and size, depending on the type, the values will be zero value, arrays start at 0
	fmt.Println(x)
	x[3] = 42 // you can replace values by index in go arrays
	fmt.Println(x)
}