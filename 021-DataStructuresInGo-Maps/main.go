package main

import "fmt"

func main() {
	// maps are like python dictionaries, an unordered list
	// of key-values mappings

	// example below of map of type map[string]int

	m := map[string]int{
		"en": 100,
		"es": 400, // a trailing comma is required!
	}

	fmt.Println(m) // prints map[en:100 es:400]
	fmt.Println(m["en"]) // prints 100
	fmt.Println(m["fr"]) // prints 0 - key doesn't exist but gets a zero value, this is very risky! you can check by...

	v, ok := m["fr"]
	fmt.Println(v, ok) // will print 0, false

	// comma ok idioms are great
	if v, ok := m["fr"]; ok {
		fmt.Println(v, "value exists")
	} else {
		fmt.Println("value doesn't exist")
	}
}