package main

import "fmt"

/*
Create a for loop using this syntax
for condition { }

Have it print out the years you have been alive.

*/

func main() {
	by := 1991
	for by <= 2021 {
		fmt.Println(by)
		by++
	}
}
