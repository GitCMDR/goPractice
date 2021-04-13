package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}


// this is how you define functions
// func (r receiver) identifier(parameters) (return(s)) {...}

// function methods are therefore defined as...

func (s secretAgent) speak(st string)  {
	fmt.Println(st)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak("Hey there, I'm")
	sa1.speak(sa1.first)
	sa1.speak(sa1.last)
}