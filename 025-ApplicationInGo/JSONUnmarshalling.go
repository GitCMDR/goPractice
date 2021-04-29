package main

import (
	"encoding/json"
	"fmt"
)

type personv2 struct { // u need an struct to host your json
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
} // cool page to do this https://mholt.github.io/json-to-go/


func main()  {
	s:= `[{"First":"L","Last":"B","Age":27},{"First":"D","Last":"S","Age":21}]` // this is my text json (raw string)
	bs := []byte(s) // i'll put my raw string into a slice of bytes

	var people []personv2 // as i have two or more objects in my json, i'll create a slice of personv2, if not, I'll get json: cannot unmarshal array into Go value of type main.personv2

	err2 := json.Unmarshal(bs, &people) // unmarshal is in place, this is why we pass the address of the value that's going to store the un-marshaled objects

	if err2 != nil { // we do a basic error check because if theres an error unmarshal will return an error
		fmt.Println(err2)
	}

	fmt.Println(people) // lets print the now un-marshaled json
}
