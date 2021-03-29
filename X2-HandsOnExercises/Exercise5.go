package main

import "fmt"

/*Create a variable of type string using a raw string literal. Print it.*/

func main() {
	myMessage := `Hello! I'm a raw-string!
                  I can do this, and it's pretty
                  cool!`

	fmt.Println(myMessage)
}
