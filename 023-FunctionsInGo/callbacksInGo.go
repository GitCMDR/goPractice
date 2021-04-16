package main

import "fmt"

// a callback is passing a func as an argument

func main() {
	fmt.Println(sum([]int{1,2,3,4,5,6,7,8,9}...))

	fmt.Println(sumEven(sum, []int{1,2,3,4,5,6,7,8,9}...))

	fmt.Println(sumOdd(sum, []int{1,2,3,4,5,6,7,8,9}...))


}

func sum(x2 ...int) int {
	total := 0
	for _, v := range x2 {
		total += v
	}
	return total
}

func sumEven(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v % 2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

// create a function that adds odd numbers

func sumOdd(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v % 2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}