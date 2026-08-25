package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func main() {
	// Crear varios usuarios e imprimirlos con distintos formatos de fmt.Printf
	var pedro User
	pedro.ID = 2
	pedro.Name = "Pedro"
	pedro.Email = "pedro@mail.com"
	pedro.Age = 42

	juan := User{
		ID:    1,
		Name:  "Juan",
		Email: "juan@mail.com",
		Age:   35,
	}

	ana := User{
		ID:   3,
		Name: "Ana",
	}

	fmt.Printf("%v\n", juan)
	fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n",
		pedro.ID, pedro.Name, pedro.Email, pedro.Age)
	fmt.Printf("{\"ID\": %d, \"Name\": \"%s\"}\n", ana.ID, ana.Name)

}
