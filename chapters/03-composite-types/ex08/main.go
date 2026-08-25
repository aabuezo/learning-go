package main

import "fmt"

func main() {
	s := []int{}

	fmt.Printf("%5s%6s%5s%5s\n", "iter", "value", "len", "cap")
	for i := 0; i < 20; i++ {
		s = append(s, i+1)
		fmt.Printf("%5d%6d%5d%5d\n", i+1, s[i], len(s), cap(s))
	}
}
