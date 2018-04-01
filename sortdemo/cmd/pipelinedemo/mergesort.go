package main

import (
	"gointro/sortdemo/pipeline"
	"fmt"
)

func main() {
	arr1 := []int {3, 2, 5, 7, 4, 6}
	arr2 := []int {3, 7, 0, 6, 5, 16}
	p := pipeline.Merge(
		pipeline.InMemorySort(pipeline.ArraySort(arr1)),
		pipeline.InMemorySort(pipeline.ArraySort(arr2)))
	for {
		if num , ok := <- p; ok {
			fmt.Println(num)
		}
	}
}
