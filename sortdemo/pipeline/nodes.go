package pipeline

import "sort"

func ArraySort(arr []int) chan int {
	out := make(chan int)

		go func() {
			for _, v := range arr {
				out <- v
			}
			close(out)
		}()

	return out
}

func InMemorySort(in chan int) chan int {
	out := make(chan int)
	go func() {
		//Read into memory
		var a []int
		for v := range in {
			a = append(a, v)
		}
		//sort a
		sort.Ints(a)

		//output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}