package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateCopy(user User, name string) {
	user.Name = name
	fmt.Println("updateCopy:", user)
}

func updateGood(user *User, name string) {
	user.Name = name
	fmt.Println("updateGood", *user)
}

func main() {
	u := User{
		Name: "john",
		Age:  43,
	}
	fmt.Println("main:", u)

	updateCopy(u, "jane")
	fmt.Println("main:", u)

	updateGood(&u, "lali")
	fmt.Println("main:", u)
}
