package main

import "fmt"

func main() {
	name := "Luis"

	tpl := `<h1>` + name + `</h1>`

	fmt.Println(tpl)
}
