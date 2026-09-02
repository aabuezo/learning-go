package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	previousPercent := debug.SetGCPercent(100)
	previousLimit := debug.SetMemoryLimit(64 << 20)

	fmt.Println("GOGC anterior:", previousPercent)
	fmt.Println("GOGC nuevo:", 100)
	fmt.Println("límite de memoria anterior:", previousLimit)
	fmt.Println("límite de memoria nuevo:", 64<<20)

	// TODO: agregar una carga pequeña y observar el comportamiento con GOMEMLIMIT.
}
