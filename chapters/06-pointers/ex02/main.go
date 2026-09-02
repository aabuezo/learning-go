package main

import "fmt"

// increment recibe la dirección de i y modifica el valor original.
func increment(i *int) {
	if i == nil {
		return
	}
	*i++
}

// decrement recibe una copia de i; su modificación se pierde al terminar.
func decrement(i int) {
	i--
}

func main() {
	i := 42
	fmt.Println(i) // 42

	increment(&i) // modifica la variable i

	fmt.Println(i) // 43

	decrement(i) // no modifica la variable i

	fmt.Println(i) // 43

}
