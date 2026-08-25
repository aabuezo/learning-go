package main

import (
	"fmt"
	"log"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Users []User

func (u Users) FindByID(id int) (*User, error) {
	for i := range u {
		if u[i].ID == id {
			return &u[i], nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

func (u Users) RemoveByID(id int) Users {
	users := make(Users, 0, len(u))
	for _, user := range u {
		if user.ID != id {
			users = append(users, user)
		}
	}
	return users
}

func main() {
	// Crear varios usuarios
	pedro := User{
		ID:    2,
		Name:  "Pedro",
		Email: "pedro@mail.com",
		Age:   42,
	}

	juan := User{
		ID:    1,
		Name:  "Juan",
		Email: "juan@mail.com",
		Age:   35,
	}

	ana := User{
		ID:    3,
		Name:  "Ana",
		Email: "ana@mail.com",
		Age:   23,
	}

	lola := User{
		ID:    2,
		Name:  "Lola",
		Email: "lola@mail.com",
		Age:   11,
	}

	// Crear un slice de User.
	users := Users{}

	// agregar usuarios con append
	users = append(users, juan)
	users = append(users, pedro, ana, lola)
	fmt.Println(users)

	// buscar un ID inexistente
	user, err := users.FindByID(5)
	if err != nil {
		fmt.Println(err)
	}
	// actualizar el email de un usuario inexistente;
	// no se puede modificar el email de un usuario inexistente
	// porque FindByID devuelve error si el ID no existe

	// buscar usuario por ID
	user, err = users.FindByID(1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*user)

	// actualizar email de un usuario
	user.Email = "user@mail.com"
	fmt.Println(*user)

	// eliminar un ID inexistente
	users = users.RemoveByID(5)

	// eliminar usuario por ID
	// verificar qué ocurre si hay dos usuarios con el mismo ID
	users = users.RemoveByID(2)
	fmt.Println(users) // se remueven todos los usuarios

}
