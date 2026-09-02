package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateUser(u User, name string) {
	u.Name = name
}

func main() {
	u := User{Name: "John", Age: 23}
	pu := &u
	fmt.Println(u)

	pu.Name = "Jane" // syntactic sugar para (*pu).Name
	fmt.Println(u)

	(*pu).Name = "Patrick"
	fmt.Println(u)

	updateUser(u, "John")
	fmt.Println(u) // {Patrick 23}
}
