package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50, 60}

	// primeros 3 elementos
	fmt.Println(nums[:3])

	// ultimos 3 elementos
	fmt.Println(nums[3:])

	// elementos del medio
	fmt.Println(nums[2:4])

	// slice vacio desde nums[:0]
	vacio := nums[:0]
	fmt.Println(vacio, len(vacio), cap(vacio))

	// slice completo con nums[:]
	completo := nums[:]
	fmt.Println(completo, len(completo), cap(completo))

	nums2 := nums[:3]
	nums2[0] = 100
	fmt.Println(nums)
	fmt.Println(nums2)
}
