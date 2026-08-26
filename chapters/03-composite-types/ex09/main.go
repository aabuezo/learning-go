package main

import "fmt"

// solucion mas limpia
func removeAt(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return nil
	}
	s := make([]int, 0)
	for i, v := range nums {
		if i == index {
			continue
		}
		s = append(s, v)
	}
	return s
}

// solucion mas eficiente porque preasigna el tamanio exacto
func removeAt2(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return nil
	}
	s := make([]int, len(nums)-1)
	si := 0
	for i, v := range nums {
		if i == index {
			continue
		}
		s[si] = v
		si++
	}
	return s
}

func main() {
	nums := []int{10, 20, 30, 40, 50}

	// eliminar primero
	newNums := removeAt(nums, 0)
	fmt.Println(newNums)

	// eliminar del medio
	newNums = removeAt(nums, len(nums)/2)
	fmt.Println(newNums)

	// eliminar ultimo
	newNums = removeAt(nums, len(nums)-1)
	fmt.Println(newNums)

	// indice invalido
	newNums = removeAt(nums, len(nums))
	fmt.Println(newNums)

	// eliminar primero
	newNums = removeAt2(nums, 0)
	fmt.Println(newNums)

	// eliminar del medio
	newNums = removeAt2(nums, len(nums)/2)
	fmt.Println(newNums)

	// eliminar ultimo
	newNums = removeAt2(nums, len(nums)-1)
	fmt.Println(newNums)

	// indice invalido
	newNums = removeAt2(nums, len(nums))
	fmt.Println(newNums)

}
