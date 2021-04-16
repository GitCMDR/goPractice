package main

import "fmt"

/*
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in
 */

func main(){
	fmt.Println(foo2([]int{1,2,3,4,5,6,7}...))
	fmt.Println(bar2([]int{1,2,3,4,5,6,7}))
}

func foo2(a ...int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

func bar2(xi []int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}