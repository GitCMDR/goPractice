package main

import "fmt"

/*Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
	1. 42
	2. “James Bond”
	3. true
  Now print the values stored in those variables using a
	1. single print statement
	2. multiple print statements
*/

func main() {
	a := 42
	b := "James Bond"
	c := true

	fmt.Printf("a = %[1]v, b = %[2]v, c = %[3]v\n",a,b,c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}