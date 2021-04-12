package main

import "fmt"

/*
Hands-on exercise #9
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

 */

func main() {
	m := map[string][]string {
		"bond_james": {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr": {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["Luis"] = []string{"David", "Noby", "Ellie"}

	for k, v := range m {
		fmt.Println(k)
		for i, j := range v {
			fmt.Println(i, j)
		}
	}

}