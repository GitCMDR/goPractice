package main

import "fmt"

/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/

func main() {
	var favSport string
	favSport = "Rugby"

	switch favSport {
	case "Football", "Volleyball", "Karate":
		fmt.Println("Favourite sport is football, volleyball or karate!")
	case "Rugby", "Wrestling":
		fmt.Println("Go Convicts, Go Silver Backs!!!")
	}
}