package main

import (
	"fmt"
)

func main() {
	x := 15
	a := &x // memory address
	fmt.Println(a)
	fmt.Println(*a)

	*a = 9

	fmt.Println(x)

}
