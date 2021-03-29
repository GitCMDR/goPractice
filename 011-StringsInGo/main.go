package main

import "fmt"

func main()  {
	s := "Hellooo, this is my string"

	bs := []byte(s)

	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
}