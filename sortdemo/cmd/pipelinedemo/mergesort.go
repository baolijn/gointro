package main

import (
	"gointro/sortdemo/pipeline"
	"fmt"
)

func main() {
	arr := []int {3, 2, 5, 7, 4, 6}
	p := pipeline.InMemorySort(pipeline.ArraySort(arr))
	for {
		if num , ok := <- p; ok {
			fmt.Println(num)
		}
	}
}
