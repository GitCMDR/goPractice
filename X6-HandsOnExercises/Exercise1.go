package main

import "fmt"

/*
Hands on exercise
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both func
print out their results
/
 */

func main() {
	fooValue := foo()
	barValueInt, barValueString := bar()

	fmt.Printf("Foo stores %v, barValue stores an int which is %v, and also a string which is %v\n", fooValue, barValueInt, barValueString)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "uno"
}
