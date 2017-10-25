package main

import (
	"fmt"
)

func main() {
	fmt.Println("HEllo")

	defer first()
	second()
}

func first() {
	fmt.Println("I am the first function")
}

func second() {
	fmt.Println("I am the second function")
}
