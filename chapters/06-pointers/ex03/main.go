package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) String() string {
	return "new user..."
}

func main() {
	p := new(int)
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

	// new(User) devuelve un puntero a un User con sus campos en el valor cero.
	pu := new(User)
	fmt.Println(pu == nil) // false
	fmt.Println(pu)        // llama a String()
	fmt.Println(*pu)       // imprime el zero value

	// &User{} también devuelve un puntero a un User con el valor cero.
	pu = &User{}
	fmt.Println(pu == nil) // false
	fmt.Println(pu)        // llama a String()
	fmt.Println(*pu)       // imprime el zero value

	// User{} usa el valor cero directamente, sin necesitar un puntero.
	u := User{}
	fmt.Println(u)  // imprime el zero value
	fmt.Println(&u) // llama a String() porque usa un puntero

	// Un literal permite crear un User con valores concretos.
	uConDatos := User{Name: "Ana", Age: 30}
	fmt.Println(uConDatos)

}
