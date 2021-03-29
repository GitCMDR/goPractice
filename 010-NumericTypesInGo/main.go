package main

import (
	"fmt"
	"runtime"
)

func main()  {
	x := 12 // if not specified, go would declare var as int
	y := 4.20 // and float64

	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
