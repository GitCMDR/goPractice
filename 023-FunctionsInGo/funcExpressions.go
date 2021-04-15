package main

import "fmt"

func main() {

	f := func() { // you can also define functions as variables
		fmt.Println("my first func expression")
	}

	f()

	y := func(x int) { // you can also define functions as variables
		fmt.Println("my second func expression", x)
	}

	y(10)

}
