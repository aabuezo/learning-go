package main

import "fmt"

type Operacion func(int, int) int

func suma(a, b int) int {
	return a + b
}

func aplicarFuncASlices(s1 []int, s2 []int, f Operacion) ([]int, error) {
	if len(s1) == 0 || len(s2) == 0 {
		return nil, fmt.Errorf("los slices no pueden estar vacios")
	}

	minLen := min(len(s1), len(s2)) // procesa minLen numeros
	lst := make([]int, 0, minLen)
	for i := 0; i < minLen; i++ {
		val := f(s1[i], s2[i])
		lst = append(lst, val)
	}

	return lst, nil
}

func main() {
	var f Operacion
	f = suma

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	res, err := aplicarFuncASlices(s1, s2, f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	f2 := func(a, b int) int {
		return a - b
	}
	res, err = aplicarFuncASlices(s2, s1, f2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	a, b := 10, 5
	fmt.Println(f(a, b))
	fmt.Println(suma(a, b))
}
