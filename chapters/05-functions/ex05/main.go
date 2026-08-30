package main

import "fmt"

func ops(a, b int) (int, int, int, error) {
	if b == 0 {
		return 0, 0, 0, fmt.Errorf("b no puede ser cero")
	}
	return a + b, a - b, a / b, nil
}

func main() {
	a, b := 10, 5
	suma, resta, division, err := ops(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("suma:", suma)
	fmt.Println("resta:", resta)
	fmt.Println("division:", division)

	a, b = 10, 0
	_, _, _, err = ops(a, b)
	if err != nil {
		fmt.Println(err)
	}
}
