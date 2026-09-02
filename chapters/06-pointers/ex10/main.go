package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	previousPercent := debug.SetGCPercent(100)
	previousLimit := debug.SetMemoryLimit(64 << 20)

	fmt.Println("GOGC anterior:", previousPercent)
	fmt.Println("GOGC nuevo:", 100)
	fmt.Println("límite de memoria anterior:", previousLimit)
	fmt.Println("límite de memoria nuevo:", 64<<20)

	const (
		blocks    = 8
		blockSize = 1 << 20
	)
	load := make([][]byte, blocks)
	for i := range load {
		load[i] = make([]byte, blockSize)
		for page := 0; page < len(load[i]); page += 4096 {
			load[i][page] = byte(i)
		}
	}

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("carga retenida: %d MiB; memoria en uso por Go: %.2f MiB; GC: %d ciclos\n",
		blocks, float64(stats.Alloc)/(1<<20), stats.NumGC)
}

// % go run chapters/06-pointers/ex10/main.go
// GOGC anterior: 100
// GOGC nuevo: 100
// límite de memoria anterior: 9223372036854775807
// límite de memoria nuevo: 67108864
// carga retenida: 8 MiB; memoria en uso por Go: 8.31 MiB; GC: 1 ciclos
