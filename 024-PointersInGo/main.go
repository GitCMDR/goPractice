package main

import "fmt"

func main() {
	a := 10
	fmt.Println(a)
	fmt.Println(&a)  // & shows you the memory address of var

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a

	// now we can get the value stored at that address using *
	fmt.Println(*b)

	fmt.Println(*&a) // & gives you the mem address
					 // * gives you the value stored at mem address
	foo(&a)
}

func foo (y *int) {
	fmt.Println(y)
	*y = 13
	fmt.Println(y)
	fmt.Println(*y)

}