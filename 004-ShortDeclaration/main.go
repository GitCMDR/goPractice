package main

import "fmt"

// identifiers name program entities such as variables and types. identifiers must always start with a letter

func main() {

	x := 42 // short declare var & assign a value
	fmt.Println(x)
	// the statement will print 42 to the console
	x = 99 // assign a new value to the already declared var
	fmt.Println(x)
	// the statement will print 99 to the console
	x = x + 100 // the operator + will sum the two operands x and 99, output will then be assigned to x
	fmt.Println(x)
	// the statement will now print 199

}
