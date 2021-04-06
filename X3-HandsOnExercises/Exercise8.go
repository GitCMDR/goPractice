package main

import "fmt"

/*
Create a program that uses a switch statement with no switch expression specified.
 */

func main() {

	switch {
	case true:
		fmt.Printf("this case is %v\n", true)
	case false:
		fmt.Printf("this case is %v\n", false)
	default:
		fmt.Println("this is my default statement :P")
	}

}
