package main

import "fmt"

/*
Hands-on exercise #2
Print every rune code point of the uppercase alphabet three times. Your output should look like this:
65
	U+0041 'A'
	U+0041 'A'
U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
 … through the rest of the alphabet characters
*/

func main() {
	for i := 1; i <= 90; i++ {
		fmt.Println(i)
		fmt.Printf("\t%[1]U '%[2]c'\n\t%[1]U '%[2]c'\n\t%[1]U '%[2]c'\n", i, i)
	}
}
