package main

import "fmt"

func main() {

	arr := [5]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for i := range arr {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
