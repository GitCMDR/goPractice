package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1,63,3622,1235,123,52,67,24,41,56}
	xs := []string{"yo","tu","el","ella","ellos","we","ze"}

	fmt.Println(xi, xs) // prints [1 63 3622 1235 123 52 67 24 41 56] [yo tu el ella ellos we ze]


	sort.Ints(xi) // sorts in place
	sort.Strings(xs) // also sorts in place

	fmt.Println(xi, xs) // prints [1 24 41 52 56 63 67 123 1235 3622] [el ella ellos tu we yo ze]

}