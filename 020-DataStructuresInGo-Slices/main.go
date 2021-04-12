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

	// deleting from slices looks more complicated than it is in python

	// Unlike other languages, Go does not provide any built-in functions to remove an element from a slice. Items need to be removed from a slice by slicing them out.

	// x is [11 22 33 44 55 66 77 88 99 88 77]

	// say i want to remove 55 from the array, i'd do
	x = append(x[:4], x[5:]...)
	fmt.Println(x)

	// x is now [11 22 33 44 66 77 88 99 88 77]

	// slices are built on top of arrays, changing slices make the compiler build a new array and drop the old one, this is another operation. if you know the specific size of your slice, and you know that it won't be changed, you could use the make builtin function to spare the compiler from running extra operations

	y := make([]int, 10, 100) // make a slice of composite literal type []int, choose a length (the number of stored values after being initialized) and then choose a capacity (The capacity of a slice is the number of elements it can hold.)

	fmt.Println(y, len(y), cap(y))
	// prints [0 0 0 0 0 0 0 0 0 0] 10 100
	// capacity doesnt mean number of indexes
	// for an index to be accessible, you'll first have to create it via append
	// once capacity is reached, the compiler will throw out the first underlying array and create a new one, doubled in size

	// you can also do multidimensional slices in go
	// you can see this by using the [][]type

	headers := []string{"locale","body_1","body_2"}
	indexes := []string{"en","hello","world"}

	df := [][]string{headers, indexes}
	fmt.Println(df)

	// lastly you can also iterate on slices

	for i, v := range headers {
		fmt.Println(i, v)
	}

}
