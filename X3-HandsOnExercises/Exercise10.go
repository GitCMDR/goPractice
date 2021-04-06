package main

import "fmt"

/*
Write down what these print:
fmt.Println(true && true)
fmt.Println(true && false)
fmt.Println(true || true)
fmt.Println(true || false)
fmt.Println(!true)
 */

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("------------------")
	/// should print
	fmt.Println(true)
	fmt.Println(false)
	fmt.Println(true)
	fmt.Println(true)
	fmt.Println(false)
}
