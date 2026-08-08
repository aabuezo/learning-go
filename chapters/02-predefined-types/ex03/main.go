package main

import (
	"fmt"
)

func main() {
	s := "Hola, mundo 🌎"
	fmt.Printf("len(s): %d\n", len(s))
	fmt.Printf("s: %s\n", s)

	for i, v := range s {
		fmt.Printf("i: %v(type: %T),\tv:%v(type: %T)\tl: %c\n", i, i, v, v, v)
	}
	fmt.Println()

	for i := 0; i < len(s); i++ {
		fmt.Printf("i: %d byte: %v char: %c\n", i, s[i], s[i])
	}

	bs := []byte(s)
	fmt.Println(bs)

	s2 := string(bs)
	fmt.Println(s2)

	fmt.Println(s == s2) // se pueden comparar con == !

	rs := []rune(s)
	fmt.Println(rs)
	fmt.Println(len(rs))
}
