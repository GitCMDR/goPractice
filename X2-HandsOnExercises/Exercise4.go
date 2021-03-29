package main

import "fmt"

/*
Write a program that
assigns an int to a variable
prints that int in decimal, binary, and hex
shifts the bits of that int over 1 position to the left, and assigns that to a variable
prints that variable in decimal, binary, and hex
 */
var x int
func main()  {
	x = 214
	fmt.Printf("%d, %b, %X\n", x, x, x)

	y := x << 1
	fmt.Printf("%d, %b, %X\n", y, y, y)
}