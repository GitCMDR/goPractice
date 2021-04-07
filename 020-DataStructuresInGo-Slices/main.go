package main

import "fmt"

func main() {
	// declaring slices looks like this in go:
	// x := []type{values}

	x := []int{11, 22, 33, 44, 55, 66, 77}

	fmt.Println(x[2]) // you can call values by index too

	// a SLICE (a type of composite literal) allows you to group values of the SAME datatype only

	// you can also loop over the slice

	for index, value := range x {
		fmt.Printf("My value is %[1]v and I can be found at Index %[2]v\n", value, index)

	}

	// you can also slice an slice! similar to python

	fmt.Println(x[1:]) // give me everything except for index 0
	fmt.Println(x[:(len(x) -1)]) // print everything except for last value


	// we can also append to slices!
	heyIReallyNeedToAppendThisToSlice := 88
	heyIHaveASliceNow := []int{99,88,77}
	x = append(x, heyIReallyNeedToAppendThisToSlice) // slice is returned, it is not an in-place change
	fmt.Println(x)

	x = append(x, heyIHaveASliceNow...) // u need to use this syntax to append a slice into another slice.
	fmt.Println(x)


}
