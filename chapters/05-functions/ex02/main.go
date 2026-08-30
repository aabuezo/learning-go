package main

import (
	"fmt"
)

func add(lst []int, val int) {
	lst = append(lst, val)
}

func addRet(lst []int, val int) []int {
	lst = append(lst, val)
	return lst
}

func main() {
	vals := []int{}

	for i := range 3 {
		add(vals, i+1)
		fmt.Printf("%v, len: %d, cap: %d\n", vals, len(vals), cap(vals))
	}

	for i := range 3 {
		vals = addRet(vals, i+1)
		fmt.Printf("vals: %v, len: %d, cap: %d\n", vals, len(vals), cap(vals))
	}

	vals = make([]int, 0, 10)

	for i := range 3 {
		add(vals, i+1)
		fmt.Printf("%v, len: %d, cap: %d\n", vals, len(vals), cap(vals))
	}

	for i := range 3 {
		vals = addRet(vals, i+1)
		fmt.Printf("vals: %v, len: %d, cap: %d\n", vals, len(vals), cap(vals))
	}
}
