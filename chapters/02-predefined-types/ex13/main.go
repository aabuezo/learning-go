package main

import "fmt"

func main() {

	m := map[string]int{}
	// m := make(map[string]int)	// otra forma

	// insertar
	m["ale"] = 51
	m["marti"] = 20
	m["yael"] = 25
	m["sil"] = 53
	m["otro"] = 30

	fmt.Println(m)

	// leer
	fmt.Println("ale:", m["ale"])

	// modificar
	fmt.Println("sil:", m["sil"])
	m["sil"] = 54
	fmt.Println("sil:", m["sil"])

	// borrar con delete
	delete(m, "otro")

	// verificar existencia con value, ok
	if v, ok := m["otro"]; ok {
		fmt.Println("otro:", v)
	}

	if v, ok := m["marti"]; ok {
		fmt.Println("marti:", v)
	}

	fmt.Println(m)
}
