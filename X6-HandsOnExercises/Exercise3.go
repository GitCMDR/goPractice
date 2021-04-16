package main

import "fmt"

/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
*/

func main() {
	defer myLastFunction()

	fmt.Println("I need to run first!")
}

func myLastFunction() {
	fmt.Println("Process is finished")
}

