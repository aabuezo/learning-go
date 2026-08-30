package main

import "fmt"

func updateSlice(lst []int, idx, val int) error {
	if idx < 0 || idx >= len(lst) {
		return fmt.Errorf("invalid index: %d", idx)
	}

	lst[idx] = val
	return nil
}

func main() {
	vals := []int{1, 2, 3, 4, 5}
	fmt.Println(vals)

	if err := updateSlice(vals, -1, 10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(vals)
	}

	if err := updateSlice(vals, 0, 10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(vals)
	}

	if err := updateSlice(vals, len(vals), 60); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(vals)
	}

	if err := updateSlice(vals, len(vals)-1, 60); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(vals)
	}
}
