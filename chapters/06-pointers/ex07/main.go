package main

import "fmt"

type Storage struct {
	value *int
}

func newValue() *int {
	value := 42
	return &value
}

func storeValue(storage *Storage, value *int) {
	storage.value = value
}

func main() {
	storage := Storage{}
	value := newValue()
	storeValue(&storage, value)

	fmt.Println(*storage.value)
	// TODO: ejecutar go build -gcflags='-m=2' . y observar los escapes.
}
