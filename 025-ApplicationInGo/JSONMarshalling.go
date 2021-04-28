package main

import (
	"encoding/json"
	"fmt"
)

type person struct { // define the struct to marshall
	First string
	last string
	age int
}

func main()  {
	p1 := person{ // create random struct 1
		First: "L", // only fields with upper case will be exported
		last:  "B",
		age:   27,
	}
	p2 := person{ // create random struct 1
		First: "D",
		last:  "S",
		age:   21,
	}

	people := []person{p1, p2} // put structs in a slice of struct

	fmt.Println(people) // prints [{L B 27} {D S 21}]

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs)) // prints [{"First":"L"},{"First":"D"}], the rest of the fields are not returned because they are lowercase
}