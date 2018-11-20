package main

import "fmt"

func main () {
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//using var
	//var name string = "Gerry"
	var name = "Gerry"
	//var age = 29
	var age int32 = 29
	const isCool = true
	//isCool = true

	// Shorthand...only works within a function
	last_name := "Laracuente"


	fmt.Println(name, age, isCool, last_name)
	fmt.Printf("%T\n", isCool)
}

