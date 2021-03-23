package main

import "fmt"

/* FYI - nice documentation and new terminology “underlying type”
https://golang.org/ref/spec#Types

For this exercise
1. Create your own type. Have the underlying type be an int.
2. create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
3. in func main
		print out the value of the variable “x”
		print out the type of the variable “x”
		assign 42 to the VARIABLE “x” using the “=” OPERATOR
		print out the value of the variable “x”
*/

type missingno int

var x3 missingno

func main() {
	fmt.Println(x3)
	fmt.Printf("%T\n", x3)
	x3 = 42
	fmt.Println(x3)
}
