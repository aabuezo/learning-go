package main

import "fmt"

func main() {
	words := []string{"go", "python", "go", "java", "go", "python"}
	fmt.Println(words)

	// Crear un map que cuente cuantas veces aparece cada palabra.
	m := map[string]int{}

	for _, w := range words {
		m[w]++
	}

	// Imprimir
	for k, v := range m {
		fmt.Println(k, v)
	}
}
