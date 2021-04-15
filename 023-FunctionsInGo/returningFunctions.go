package main

import "fmt"

// you can also return functions in go!

func main() {

	e1 := foo5()
	fmt.Println(e1)

	e2 := func() int {
		return 451
	}()

	fmt.Printf("im e2 %v, %T\n", e2, e2)

	e3 := bar5()

	fmt.Printf("im e3 %v, %T\n", e3, e3)

	fmt.Println(e3())
}


func foo5() string {
	s3 := "Hello!"
	return s3
}

func bar5() func() int {
	return func() int {
		return 200
	}
}