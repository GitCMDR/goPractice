package main

import "fmt"

// functions help avoid spaghetti code
// they are very similar to python functions
// a function keyword is func
// then we have something called receiver (r receiver)
// receiver are a lot like parameters, but it's going to be attached to a type
// then we have the identifier (name of the func)
// and then we have parameters that we passed to the function (parameters)
// and then we have returns, which can also be multiple, or singular (return(s))
// all in all, a function call can look like:
// func (r receiver) identifier(parameters) (return(s)){ code... }


func main()  {

	// you can use defer to make sure your function runs
	// at the very last step of your program

	defer runMeLast() // will run at the end

	foo()
	bar("Noby")
	fmt.Println(woo("Ellie"))
	catName, isBlack := cat("Misifus", "Black")

	fmt.Printf("The cat is called %v and it's %v that its colour is black\n", catName, isBlack)

	myFoobarsum := foobar(1,2,3,4,5,6,7,8,9,0)
	fmt.Println(myFoobarsum)

	// myFoobarsum = foobar([]int{1,3,5,7,9}) this wouldn't work at all, yes, our function expects an unlimited number of ints, and yes these ints get grouped together as an int slice inside the function, but an slice of ints is NOT a number of ints. We can easily do this conversion using unfurling.

	myFoobarsum = foobar([]int{1,3,5,7,9}...) // ... is magic!
	fmt.Println(myFoobarsum)

	// with variadic parameters, you can also pass zero parameters, because an unlimited number of passed ints (in this specific example) also covers ZERO passed ints. aka, a variadic parameter takes into account from ZERO to INFINITE.

	myFoobarsum = foobar()
	fmt.Println(myFoobarsum)

}

func foo() {
	fmt.Println("Hello there! woof")
} // function that calls statements directly

func runMeLast() {
	fmt.Println("Please run me last")
}

func bar(s string) {
	fmt.Printf("Hello %v\n", s)
} // function that uses parameters as part of its statement

func woo(s string) string {
	return fmt.Sprint("Hello there ", s)
} // function that uses parameters as part of its statement and also returns something (singular return)

func cat(name string, colour string) (string, bool){

	var isBlack bool

	if colour == "Black" {
		isBlack = true
	} else {
		isBlack = false
	}
	return name, isBlack
} // you can pass multiple values to a function, and also return more than a singular value

// you can also pass variadic parameters to a function, and it looks like. VARIADIC parameters can only be passed as THE LAST element of your parameters in the function

func foobar(x ...int) int { // this is expecting an unlimited number of ints, NOT A SLICE OF INTS
	fmt.Println(x)
	fmt.Printf("%T\n", x) // x is then processed as an int slice

	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}