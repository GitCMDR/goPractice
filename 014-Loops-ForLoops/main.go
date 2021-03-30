package main

import "fmt"

func main() {
	x := 1
	for x < 10 { // this is a for statement in go with a single condition
		fmt.Println(x)
		x++
	}
	fmt.Println("Loop is finished")
}