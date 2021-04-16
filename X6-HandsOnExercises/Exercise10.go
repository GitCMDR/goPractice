package main

import "fmt"

/*
Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:
 */

func main()  {

	myEnclosedVariable := 20

	fmt.Printf("Enclosed variable in main is %v\n", myEnclosedVariable)

	someFunction()
}

func someFunction() {
	myEnclosedVariable := 23
	fmt.Printf("Enclosed value inside function is %v\n", myEnclosedVariable)
}
