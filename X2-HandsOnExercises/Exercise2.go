package main

import "fmt"

/* Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
*/

func main()  {
	a := 1 == 3
	b := 2 <= 4
	c := 5 >= 4
	d := 1000 != 10000
	e := 4 < 10
	f := 2 > 100

	fmt.Print(a,b,c,d,e,f,"\n")
}
