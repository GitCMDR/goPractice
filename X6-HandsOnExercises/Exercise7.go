package main

/*
Assign a func to a variable, then call that func
 */

import "fmt"

func main() {
	f := func() {fmt.Println("hello")}
	f()
}
