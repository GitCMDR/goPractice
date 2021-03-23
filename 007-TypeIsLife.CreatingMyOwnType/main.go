package main

import "fmt"

var a int
var unauthentic string

type queen string // define a new type called queen, that copies string

var b queen // define a queen variable

func main()  {
	a = 10
	b = "RuPaul" // RuPaul is a queen, that copies string
	unauthentic = "Silky" // Someone is a normal string
	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)

	// b = unauthentic // This will error out as my variable of type queen is no longer a string, as they are technically different types, you can't overwrite one with the other

	// we can however use conversion to convert unauthentic into a queen
	b = queen(unauthentic)
	fmt.Printf("%v is of type %T\n", b, b)


}