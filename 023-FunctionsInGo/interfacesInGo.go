package main

import "fmt"

type person struct {
	first string
	last string
}

// type = keyword, human = identifier, interface = type
type human interface {
	speak()
	// what we are defining with this is,
	// anyone who has the method speak is also the type of human, this is what interface is all about!
}

func (p person) speak()  {
	fmt.Println("Hi im person")
}

// A VALUE CAN BE OF MORE THAN ONE TYPEEEEEE IN GO WOW

func bar2(h human) {
	fmt.Println("I'm the bar2 function called via polymorphism on", h)
}

func main() {
	p1 := person{
		first: "Dr.",
		last: "Patiño",
	}

	bar2(p1)
}