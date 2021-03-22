package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	foo()
	fmt.Println("Print something else")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm calling Foo from within main")
}

// control flow:
// (1) sequential
// (2) loop, iterative
// (3) conditional
