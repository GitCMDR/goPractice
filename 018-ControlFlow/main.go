package main

import "fmt"

func main()  {
	switch {
	case true:
		fmt.Println("This was false!")
	//case (2 == 4):
		fmt.Println("2 is not equal to 4")
	//case true:
		fmt.Println("This was true!")
		fallthrough
		// you can also add a fallthrough to get the next condition to run
	case false:
		fmt.Println("a is a!")
	//case true:
		fmt.Println("b is b! will the fallthrough get me?, no unless the fallthrough is directly above this case")
		// if nothing is true, then you can also specify a default in go
		default:
			fmt.Println("this is the default!")
	}
	fmt.Println()
	// you can also switch on a value
	switch "Bond" {
	case "James":
		fmt.Println("I'm James")
	case "Peter":
		fmt.Println("I'm Peter!")
	default:
		fmt.Println("I couldn't find Bond!")
	}
	fmt.Println()
	// you can also switch on a value inside a variable
	caseVariable := "Peter"
	switch caseVariable {
	case "P":
		fmt.Println("Not Bond")
	case "Peter":
		fmt.Println("I'm Peter!")
	default:
		fmt.Println("I couldn't find Bond!")
	}
	fmt.Println()
	// you can also have different cases in one case
	caseVariable2 := "Jaime"
	switch caseVariable2 {
	case "James", "Jaime", "Jimmy":
		fmt.Println("I'm James, but you can call me Jimmy!")
	case "Peter":
		fmt.Println("I'm Peter!")
	default:
		fmt.Println("I couldn't find Bond!")
	}
}