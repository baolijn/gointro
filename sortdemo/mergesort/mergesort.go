package main

import "fmt"

func merge(arr []int, l int, mid int, r int)  {
	var aux = arr
	var i = l
	var j = mid + 1
	for k:=l; k < r; k ++ {
		if l > mid {
			aux[k] = arr[r-j]
			j++
		} else if j > r {
			aux[k] = arr[i-l]
			i++
		} else if arr[i -l] < arr[r-j] {
			aux[k] = arr[i-l]
			i ++
		} else {
			aux[k] = arr[r-j]
			j ++
		}
	}
}

func mergesort(arr []int)  {
	_mergesort(arr, 0, len(arr)-1)
}

func _mergesort(arr []int, l int, r int) {
	mid := (l + r)/2;
	_mergesort(arr, l, mid)
	_mergesort(arr, mid + 1, r)
	merge(arr, l, mid, r)
}

func main() {
	a := []int {3, 6, 5, 2, 7}
	mergesort(a)
	fmt.Println(a)
}
