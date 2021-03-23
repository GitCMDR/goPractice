package main

import "fmt"

// func Println(a ...interface{}) (n int, err error)
// ... means a non-specific number of values
// interface{} means a value of any type
/// ^ together, they are a variadic parameter

func main() {
	n, err := fmt.Println("Hey")
	fmt.Println(n)
	fmt.Println(err)
}

// above would return
// Hey
// 4
// err: <nil>, err can't be thrown out with a '_'

// extra bit of knowledge: you can't have a variable and do nothing with it in go - you need to be explicit with the compiler. never have a variable that you don't use, that's bad programming.

// also: just to remember, every program has to have a package main and a func main.
