package main

import "fmt"

func main() {

	age := 20

	switch age {
	case 10:
		fmt.Println("Watch pokemon")
	case 15:
		fmt.Println("Watch Dragonball Z")
	case 20:
		fmt.Println("Watch Naruto")
	default:
		fmt.Println("Watch what you feel like")
	}
}
