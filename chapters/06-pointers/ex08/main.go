package main

import (
	"fmt"
	"runtime"
)

func printStats(label string) {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("%s: Alloc=%d HeapAlloc=%d HeapObjects=%d NumGC=%d\n",
		label, stats.Alloc, stats.HeapAlloc, stats.HeapObjects, stats.NumGC)
}

func main() {
	printStats("antes")

	data := make([][]byte, 0, 1000)
	for i := 0; i < 1000; i++ {
		data = append(data, make([]byte, 1024))
	}
	printStats("después de asignar")

	data = data[:10]
	runtime.GC()
	printStats("después de GC")
	_ = data
}
