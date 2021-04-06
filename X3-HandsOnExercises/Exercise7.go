package main

import "fmt"

/*
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
 */

func main() {
	if 0 != 1 {
		fmt.Println("0 is not 1")
	} else if 0 == 0 {
		fmt.Println("0 is 0")
	} else {
		fmt.Println("We're dealing in quantum now")
	}
}