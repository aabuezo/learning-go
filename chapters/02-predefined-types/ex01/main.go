package main

import (
	"fmt"
	"math"
	"unsafe"
)

var (
	i   int
	i64 int64
	i32 int32
	f64 float64
	b   byte
	r   rune
	s   string
	t   bool
)

func main() {
	i = 10
	// i64 = i

	i32 = int32(i)
	i64 = int64(i32)
	// i32 = i64

	f64 = float64(i64)
	fmt.Printf("f64: %.2f, %T\n", f64, f64)

	b = byte(i)
	f64 = 3.14
	f32 := float32(f64)
	fmt.Printf("f32: %.6f, %T\n", f32, f32)

	b = 65
	r = rune(b)
	s = string(r)
	fmt.Printf("s: %s, %T\n", s, s)

	r = '🌎'
	fmt.Printf("r: %c, %T\n", r, r)

	// t = bool(i) //cannot convert i (variable of type int) to type bool
	t = f64 > float64(i)
	// s = string(t)
	fmt.Printf("t: %v, %T\n", t, t)

	// b = byte(300) // constant 300 overflows byte
	// fmt.Printf("b: %v\n", b)

	fmt.Println("int8:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, math.MaxInt64)

	fmt.Println("uint8:", 0, math.MaxUint8)
	fmt.Println("uint16:", 0, math.MaxUint16)
	fmt.Println("uint32:", 0, math.MaxUint32)
	fmt.Println("uint64:", uint64(0), uint64(math.MaxUint64))

	fmt.Println("float32 max:", math.MaxFloat32)
	fmt.Println("float64 max:", math.MaxFloat64)

	fmt.Println("int min:", math.MinInt)
	fmt.Println("int max:", math.MaxInt)

	// byte == uint8
	fmt.Println("byte:", 0, math.MaxUint8)

	// rune == int32
	fmt.Println("rune:", math.MinInt32, math.MaxInt32)

	fmt.Println("int:", unsafe.Sizeof(i), "bytes")
	fmt.Println("int32:", unsafe.Sizeof(i32), "bytes")
	fmt.Println("int64:", unsafe.Sizeof(i64), "bytes")

}
