package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	a := nums[:2]

	fmt.Println(nums)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 30)
	fmt.Println(nums)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 40)
	fmt.Println(nums)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 50)
	fmt.Println(nums)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 60) // aca ya no modifica el backing array
	fmt.Println(nums)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 70)
	fmt.Println(nums)
	fmt.Println(a, len(a), cap(a))
}
