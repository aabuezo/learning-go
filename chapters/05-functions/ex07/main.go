package main

import (
	"fmt"
)

func Open(s string) {
	fmt.Println("resource opened!", s)
}

func Close(s string) {
	fmt.Println("resource closed!", s)
}

func Call() {
	Open("from Call")
	defer Close("from Call")
	fmt.Println("Processing sth in Call...")
}

func main() {

	Open("from main")
	defer Close("from main")
	Call()
	defer func() {
		fmt.Println("2nd defer...")
	}()

	fmt.Println("in main...")
}
