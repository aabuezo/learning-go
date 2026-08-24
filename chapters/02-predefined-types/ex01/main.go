package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// default types, zero values
	var b byte        // alias para uint8
	var i int         // puede ser int32 o int64 dependiendo de la plataforma
	var ui uint       // uint32 o uint64
	var r rune        // int32
	var uiptr uintptr // entero sin signo capaz de guardar una direccion como numero

	// tipos enteros explicitos
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	// flotantes
	var f32 float32
	var f64 float64

	// constantes implicitas
	const ic = 10 // ic es untyped
	b = ic        // puede ser asignada a distintos tipos numericos
	i = ic
	ui = ic
	const dc = 3.14      // c es untyped
	var fci float32 = dc // c puede ser asignada a float32
	var dci = dc         // o a float64, declaracion implicita
	var dce float64 = dc // decl explicita

	// constantes explicitas
	const ie int = 10 // ie es una constante de tipo int
	var cie int = ie  // y solo puede ser asignada a ints
	// var ci8e int8 = ie 			// pero no a otro tipo int

	// tipo runa
	r = 'a' // Unicode codepoint

	// tipo string
	var si = "Hola, "        // decl implicita
	var se string = "mundo!" // decl explicita
	var sraw string = `
	Hello,
		world!
	`

	// tipo boolean
	var bf bool         // zero value es false
	var bt = true       // decl implicita
	var btt bool = true // decl explicita

	// short declaration
	bff := false

	// conversiones entre tipos
	b = byte(i)        // conversion explicita de int a byte
	ui64 = uint64(i)   // de int a uint64
	fci = float32(dce) // de float64 a float32
	dce = float64(fci) // de float32 a float64
	//r = sraw			// no es posible
	i = 65
	r = rune(i)                         // 'A'
	s := string(r)                      // "A"
	pi := &i                            // puntero a int
	uiptr = uintptr(unsafe.Pointer(pi)) // usar solamente en codigo de bajo nivel o interop con C

	fmt.Println(b, i, ui, r, uiptr)
	fmt.Println(i8, i16, i32, i64, ui8, ui16, ui32, ui64)
	fmt.Println(f32, f64, fci, dci, dce)
	fmt.Println(cie)
	fmt.Println(si, se, sraw)
	fmt.Println(bf, bt, btt, bff)
	fmt.Println(s)

	fmt.Printf("valor: %d\n", i)
	fmt.Printf("direccion: %#x\n", uiptr)

}
