package main

import "fmt"

/*
Create a func which returns a func
assign the returned func to a variable
call the returned func
 */

func main() {
	f2 := funky()
	f2()
}

func funky() func() {
	return func(){fmt.Println("I'm a returned func!")}
}
