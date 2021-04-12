package main

import "fmt"

// structs look very similar to js objects

type dog struct {
	name string
	fur string
	breed string
	years int
}

// there's also something called embedded structs, which are structs inside structs

type petRegistry struct {
	dog
	microchipNumber int
}

func main() {

	noby := dog {
		name: "Noby",
		fur: "Black and White",
		breed: "Staffy",
		years: 5,
	}

	ellie := dog {
		name: "Ellie",
		fur: "Brindle",
		breed: "Bull Arab",
		years: 5,
	}

	nobyRegister := petRegistry{
		dog: dog{ // this fields will get promoted, you can access them directly
			name: "Noby", // but you can call them with full dot notation
			fur: "Black and White", // e.g nobyRegister.dog.name
			breed: "Staffy", // or more easily nobyRegister.name
			years: 5,
		},
		microchipNumber: 123456,
	}

	fmt.Println("Hello ", ellie.name, "and", noby.name)
	fmt.Println(nobyRegister.name, nobyRegister.microchipNumber)

	// if you don't feel you need to define a whole struct (maybe you want to use a struct once or something) you can also do the following

	veterinary := struct {
		name string
		address string
	}{
		name: "Arcoiris",
		address: "Stairway to heaven",
	}

	fmt.Println(veterinary)
}