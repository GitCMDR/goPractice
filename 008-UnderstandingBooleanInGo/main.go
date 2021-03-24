package main

import "fmt"

// kinda basic, but a bool value can be true or false

var x bool // declare x with underlying type bool

func main() {
	a := 7
	b := 8

	fmt.Println(x) // booleans zero value is false
	x = true
	fmt.Println(x)

	/// comparisons return boolean values e.g:
	fmt.Println(a == b)
}