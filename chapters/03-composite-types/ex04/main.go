package main

import "fmt"

func main() {
	var a []int
	b := []int{}
	c := make([]int, 0)

	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(c == nil) // false

	fmt.Printf("len(a): %d, cap(a): %d, a: %v\n", len(a), cap(a), a)
	fmt.Printf("len(b): %d, cap(b): %d, b: %v\n", len(b), cap(b), b)
	fmt.Printf("len(c): %d, cap(c): %d, c: %v\n", len(c), cap(c), c)

	a = append(a, 1, 2, 3)
	b = append(b, 1, 2, 3)
	c = append(c, 1, 2, 3)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
