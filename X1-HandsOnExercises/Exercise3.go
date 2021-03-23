package main

import "fmt"

/*
	1. Using the code from the previous exercise,
	At the package level scope, assign the following values to 	the three variables

	for x assign 42
	for y assign “James Bond”
	for z assign true

	2. in func main
	2.1 use fmt.Sprintf to print all of the VALUES to one single 	string.
	2.2 ASSIGN the returned value of TYPE string using the 	short declaration operator to a VARIABLE with the IDENTIFIER “s”
	3. print out the value stored by variable “s”
*/

var x2 = 42
var y2 = "James Bond"
var z2 = true

func main() {
	s := fmt.Sprintf("%v %v %v", x2, y2, z2)
	fmt.Println(s)
}