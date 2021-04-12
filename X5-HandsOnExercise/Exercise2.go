package main

import "fmt"

/*
Hands-on exercise #2
Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
 */

type person2 struct {
	firstName string
	lastName string
	favouriteIceCreamFlavours []string
}


func main()  {

	p1 := person2 {
		firstName: "luis",
		lastName: "b",
		favouriteIceCreamFlavours: []string{"vanilla", "chocolate", "strawberry"},
	}

	p1Map := map[string]person2 {
		p1.lastName : p1,
	}

	for _,v := range p1Map {
		fmt.Println(v.firstName, v.lastName)
		for i,v2 := range v.favouriteIceCreamFlavours {
			fmt.Println(i, v2)
		}
	}
}