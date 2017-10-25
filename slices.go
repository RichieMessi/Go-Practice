package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// s := []string{"one tails", "two tails", "three tails"}
	// fmt.Println(s)

	// r := make([]string, 3)

	// // fmt.Println(r)

	// r[0] = "Hi"
	// r[1] = "Hello"
	// r[2] = "Hurray"

	// fmt.Println(r)

	fmt.Println("The square root of 4 is ", math.Sqrt(4))
	fmt.Println("The random number is  ", rand.Intn(20))

	var a = 73
	var b = float64(a)
	x := a

	fmt.Println(b)
	fmt.Println(x)

}
