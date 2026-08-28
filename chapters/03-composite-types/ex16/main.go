package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Users struct {
	users map[int]User
}

// insertar usuario
func (us *Users) Add(u User) {
	us.users[u.ID] = u
}

// buscar usuario
func (us *Users) Search(id int) (User, error) {
	// Los elementos de un map no se pueden modificar directamente
	// mediante un puntero.
	user, ok := us.users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// actualizar usuario
func (us *Users) Update(u User) error {
	_, ok := us.users[u.ID]
	if !ok {
		return errors.New("user not found")
	}
	us.users[u.ID] = u
	return nil
}

// borrar usuario
func (us *Users) Delete(id int) error {
	if _, ok := us.users[id]; ok {
		delete(us.users, id)
		return nil
	}
	return errors.New("user not found")
}

// Imprimir
func (us Users) Print() {
	for _, user := range us.users {
		fmt.Println(user)
	}
}

// Constructor
func NewUsers() Users {
	return Users{users: make(map[int]User)}
}

func main() {

	// Crear varios usuarios
	pedro := User{
		ID:    1,
		Name:  "Pedro",
		Email: "pedro@mail.com",
		Age:   42,
	}

	juan := User{
		ID:    2,
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
		ID:    4,
		Name:  "Lola",
		Email: "lola@mail.com",
		Age:   11,
	}

	users := NewUsers()

	users.Add(pedro)
	users.Add(juan)
	users.Add(ana)
	users.Add(lola)

	users.Print()
	fmt.Println("---------------------------")

	user, err := users.Search(10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("---------------------------")

	user, err = users.Search(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	fmt.Println("---------------------------")

	user.Age = 44
	if err = users.Update(user); err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	fmt.Println("---------------------------")

	if err = users.Delete(4); err != nil {
		fmt.Println(err)
	}
	users.Print()
}
