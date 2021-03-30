package main

import "fmt"

func main() {
	// for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Println("I'm i number", i)
		for j := 0; j <= 5; j++ {
			fmt.Println("I'm j number", j)
		}
	}
}
