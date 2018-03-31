package main

import (
	"sort"
	"fmt"
)

func main() {
	//create a slice of int
	a := []int {3, 6, 5, 2, 7}
	sort.Ints(a)
	fmt.Println(len(a))
	for i, v := range a {
		fmt.Println(i, v)
	}
}
