package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 3)

	fmt.Println("src:", src)

	total := copy(dst, src)
	fmt.Println("dst:", dst) // [1 2 3]
	fmt.Println("total elementos copiados:", total)

	dst = make([]int, 10) // crea un slice con 10 elementos 0
	total = copy(dst, src)
	fmt.Println("dst:", dst) // [1 2 3 4 5 0 0 0 0 0]
	fmt.Println("total elementos copiados:", total)
}
