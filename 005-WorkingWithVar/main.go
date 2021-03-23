package main

import "fmt"

// compared to var, short declaration can only works inside a function

var y = "hey"
var z int


// you can also define the type of the variable while declaring the var, this will initialize the variable with its respective zero value. zero values change depending on the data type, so an int would be 0, whereas a str would be something like "", 0.0 for floats, and nil for pointers, functions, interfaces, channels and maps

func main() {
	x := 10

	fmt.Printf("i said %v %v times\n", y, x)

	fmt.Println(z)

}

// best practice is to limit the scope of the variables though, so it's better to stick to short declarations as much as possible