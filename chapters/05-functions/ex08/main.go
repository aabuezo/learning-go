package main

func crearContador() func() int {
	count := 0
	return func() int { // closure
		count++
		return count
	}
}

func main() {

	count1 := crearContador()
	println(count1()) // 1
	println(count1()) // 2
	println(count1()) // 3

	count2 := crearContador()
	println(count2()) // 1
	println(count2()) // 2
	println(count2()) // 3
}
