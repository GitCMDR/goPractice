package main

import "fmt"

/*
Build and use an anonymous func
*/

func main() {
	func(){
		fmt.Println("Hey. I'm an anonymous function!")
	}()
}