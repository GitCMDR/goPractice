package main

import "fmt"

var y = 10
var z = "this is a nice string"
var a = `one day he said "i'll come back"` // in case I need to print quotation marks, you can start strings with back quotes

func main() {
	fmt.Println(y) // this will print 10
	fmt.Printf("%T\n", y) // this will print int, bless you printf

	// z = 10 you can't really do this, go is a statically typed language, you can do this in JS or Python, but not go
	fmt.Println(z)
	fmt.Println(a)
}