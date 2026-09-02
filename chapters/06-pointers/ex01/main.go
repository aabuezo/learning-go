package main

import "fmt"

func main() {
	i := 42
	p := &i         // p apunta a i (contiene la direccion de memoria de i)
	fmt.Println(i)  // imprime 42
	fmt.Println(*p) // imprime 42
	fmt.Println(p)  // imprime la direccion de i

	*p++ // incrementa el valor de i a través del puntero p

	fmt.Println(i)  // imprime 43
	fmt.Println(*p) // imprime 43
	fmt.Println(p)  // imprime la direccion de i
}
