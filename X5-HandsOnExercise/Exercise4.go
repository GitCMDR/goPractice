package main

import "fmt"

/*
Hands-on exercise #4
Create and use an anonymous struct.
 */

func main() {
	myStruct := struct {
		hp int
		mp int
		class string
	}{
		hp: 100,
		mp: 230,
		class: "priest",
	}

	fmt.Printf("my %v has %v HP and %v MP\n", myStruct.class, myStruct.hp, myStruct.mp)
}