package main

import "fmt"

func main() {
	ages := map[string]int{
		"Ana":  30,
		"Luis": 25,
		"Juan": 30,
	}
	// Crear un nuevo map:
	byAge := map[int][]string{}

	// Donde la clave sea la edad y el valor sea la lista de nombres con esa edad.
	for k, v := range ages {
		byAge[v] = append(byAge[v], k)
	}

	// Imprimir
	for k, v := range byAge {
		fmt.Println(k, v)
	}
}
