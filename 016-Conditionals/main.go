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


	if x := 42; x == 2 {
		fmt.Println("Hellooo!")
	}
}