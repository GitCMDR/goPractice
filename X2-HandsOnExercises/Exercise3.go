package main

import "fmt"

/*
Create TYPED and UNTYPED constants. Print the values of the constants.
 */

const (
	a = 1
	b int = 2
)

func main()  {
	fmt.Println(a, b)
}