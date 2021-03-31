package main

import "fmt"

func main()  {
	if true {
		fmt.Println("Hellooo!")
	}

	if false {
		fmt.Println("Hellooo!")
	}

	if 2 == 2 {
		fmt.Println("Hellooo!")
	}

	if 0 != 1 {
		fmt.Println("Hellooo!")
	}


	if x := 42; x == 2 { // This is an initialization if statement
		fmt.Println("Hellooo!")
	}

	// you can also do something cool like, where you che
	if x, y := 10, 30; x == 10 {
		fmt.Println(x, y)
	}

	// apart from if, we also have 'else if' 'else'

	if x := 10; x == 10 {
		fmt.Println("X is 10")
	} else if x == 20 {
		fmt.Println("X is 20")
	} else {
		fmt.Println("X is not 10 nor 20")
	}
}