package main

import (
	"fmt"
)

func main() {

	// var a int = 10
	// var b int8 = 5

	// fmt.Println(int8(a) + b)
	// fmt.Println(a + int(b)) // types should be the same or the compiler gives an error
	// No implicit conversions..
	// Golang is statically styped and strongly typed lang

	// a := 10
	// b := 3
	// fmt.Println(a & b) // 2
	// fmt.Println(a | b) // 11

	// fmt.Println(a << 4)
	// fmt.Println(a >> 3)
	// // bit conversion

	// x := 1 + 4i
	// var y complex64 = 1 + 2i
	// var z complex128 = complex(5, 12)

	// fmt.Printf("%v, %T", x, x)

	// fmt.Println("")

	// fmt.Printf("%v, %T", real(y), real(y))

	// fmt.Println("")

	// fmt.Printf("%v, %T", z, z)

	// var s string = "Willy Wonka"

	// fmt.Printf("%v, %T", s[2], s[2])

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
