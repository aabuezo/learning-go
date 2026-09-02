package main

import "runtime"

func main() {
	// Ejecutar, por ejemplo, con GOGC=25 o GOGC=200 y gctrace=1.
	for etapa := 0; etapa < 10; etapa++ {
		data := make([][]byte, 0, 10000)
		for i := 0; i < 10000; i++ {
			data = append(data, make([]byte, 1024))
		}
		runtime.KeepAlive(data)
	}
}
