package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func findUser(id int) *User {
	users := []User{
		{ID: 1, Name: "Ana"},
		{ID: 2, Name: "Luis"},
	}

	for i := range users {
		if users[i].ID == id {
			return &users[i]
		}
	}
	return nil
}

func main() {
	for _, id := range []int{1, 99} {
		user := findUser(id)
		if user == nil {
			fmt.Println("usuario no encontrado")
			continue
		}
		fmt.Println(user.Name)
	}
}
