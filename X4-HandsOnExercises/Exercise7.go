package main

import "fmt"

/*
Hands-on exercise #7
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.
 */

func main() {
	a := [][]string{[]string{"James", "Bond", "Shaken, not stirred"},[]string{"Miss", "Moneypenny", "Helloooooo, James."}}

	fmt.Println(a)

	for _, v := range a {
		fmt.Println(v)
		for _,j := range v {
			fmt.Println(j)
		}
	}



}
