package main

import "fmt"

func main() {

	arr := [3]int{1, 2, 3}

	fmt.Println("en main():", arr)

	updateArr(arr)

	fmt.Println("en main():", arr)

	updateArrPointer(&arr)

	fmt.Println("en main():", arr)
}

// no modifica al array porque usa una copia (pass by value)
func updateArr(a [3]int) {
	a[2] = 4
	fmt.Println("en updateArr():", a)
}

// para modificar el original hay que pasar un puntero
func updateArrPointer(pa *[3]int) {
	pa[2] = 4
}
