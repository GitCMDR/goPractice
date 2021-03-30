package main

import "fmt"

// Mini-challenge. Print ASCII values from 33 up to and including 122

func main() {
	for i:= 33; i < 123; i++ {
		fmt.Printf("The number %v is decimal for %q for ASCII\n", i, i)
	}
}
