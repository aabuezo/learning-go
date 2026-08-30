package main

import "fmt"

func add(m map[string]int, k string, v int) {
	m[k] = v
}

func main() {
	m := map[string]int{}
	fmt.Println(m)

	add(m, "hello", 1)
	add(m, "world", 2)
	fmt.Println(m)
}
