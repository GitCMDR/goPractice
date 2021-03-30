package main

/* Using iota, create 4 constants for the NEXT 4 years. Print the constant values. */

import "fmt"

const (
	_ = iota
	A = 2021 + iota
	B = 2021 + iota
	C = 2021 + iota
	D = 2021 + iota

)

func main() {

	fmt.Println(A,B,C,D)
}
