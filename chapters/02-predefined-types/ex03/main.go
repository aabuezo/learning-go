package main

import (
	"fmt"
)

func main() {
	// len e índices trabajan con bytes, mientras que range sobre un string recorre runes

	s := "Hola, mundo 🌎!"
	fmt.Println(s)

	// len(s)
	fmt.Println("len(s):", len(s))

	// recorrido por índice
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i]) // cada byte como hexadecimal
	}
	fmt.Println()

	// recorrido con range
	for _, r := range s {
		fmt.Printf("%c ", r) // como caracter (rune)
	}
	fmt.Println()

	for _, r := range s {
		fmt.Printf("%v ", r) // como punto de codigo Unicode (rune)
	}
	fmt.Println()

	// posición e índice de cada carácter
	for i, r := range s {
		fmt.Printf("índice: %d, rune: %c, valor: %v\n", i, r, r)
	}
}
