package main

import "fmt"

func main() {

	var a map[string]int
	b := map[string]int{}
	c := make(map[string]int)

	// a es nil
	// con un map nil se puede leer, pero no se puede insertar
	if v, ok := a["algo"]; ok {
		fmt.Println("a[\"algo\"]:", v)
	} else {
		fmt.Println("a - ok: false, v:", v)
	}

	// a["algo"] = 1	// panic: assignment to entry in nil map
	// if v, ok := a["algo"]; ok {
	// 	fmt.Println("a[\"algo\"]:", v)
	// } else {
	// 	fmt.Println("a - ok: false, v:", v)
	// }

	// b
	if v, ok := b["algo"]; ok {
		fmt.Println("b[\"algo\"]:", v)
	} else {
		fmt.Println("b - ok: false, v:", v)
	}

	b["algo"] = 1
	if v, ok := b["algo"]; ok {
		fmt.Println("b[\"algo\"]:", v)
	} else {
		fmt.Println("b - ok: false, v:", v)
	}

	// c
	if v, ok := c["algo"]; ok {
		fmt.Println("c[\"algo\"]:", v)
	} else {
		fmt.Println("c - ok: false, v:", v)
	}

	c["algo"] = 1
	if v, ok := c["algo"]; ok {
		fmt.Println("c[\"algo\"]:", v)
	} else {
		fmt.Println("c - ok: false, v:", v)
	}
}
