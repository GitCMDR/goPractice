package main

import "fmt"

// Write a program that prints a number in decimal, binary, and hex

func main() {
	a := 123

	fmt.Printf("Decimal is %[1]d, binary is %[1]b and finally, hex is %[1]X\n", a)
}