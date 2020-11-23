package main

import (
	"fmt"
)

func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	list := []int{65, 85, 74, 54, 14, 38, 67, 50, 3, 99, 79, 11, 73, 1, 60, 501, 7, 83, 30, 77, 12, 19, 52, 21, 95, 89, 4, 101, 49, 29, 45, 41, 95, 82, 105}

	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
