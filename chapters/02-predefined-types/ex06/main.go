package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [4]int{1, 2, 3, 4}
	arr4 := [4]int{1, 2, 3, 5}
	fmt.Println(arr1 == arr2) // true
	// fmt.Println(arr1 == arr3)	// invalid operation: arr1 == arr3 (mismatched types [3]int and [4]int)
	fmt.Println(arr3 == arr4) // false
}
