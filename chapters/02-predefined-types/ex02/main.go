package main

import (
	"fmt"
)

func main() {
	a := 7
	b := 2

	int_div := a / b
	fmt.Println("int_div:", int_div)

	float_div := float64(a / b) // convierte el resultado de la division entera
	fmt.Println("float_div:", float_div)

	// float_div = a / float64(b)	// invalid operation: a / float64(b) (mismatched types int and float64)
	float_div = float64(a) / float64(b)
	fmt.Println("float_div:", float_div)

	mod := a % b
	fmt.Println("mod:", mod)
}
