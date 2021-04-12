package main

import "fmt"

/*
Hands-on exercise #1
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
*/

type person struct {
	firstName string
	lastName string
	favouriteIceCreamFlavours []string
}

func main()  {

	p1 := person {
		firstName: "luis",
		lastName: "b",
		favouriteIceCreamFlavours: []string{"vanilla", "chocolate", "strawberry"},
	}

	p2 := person {
		firstName: "clau",
		lastName: "b",
		favouriteIceCreamFlavours: []string{"bubblegum", "pistachio", "kinder"},
	}

	fmt.Println(p1.firstName, p1.lastName)
	for _,v := range p1.favouriteIceCreamFlavours {
		fmt.Println(v)
	}

	fmt.Println(p2.firstName, p2.lastName)
	for _,v := range p2.favouriteIceCreamFlavours {
		fmt.Println(v)
	}

}