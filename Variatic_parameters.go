package main

import (
	"fmt"
)

func main() {

	r := add(2, 3, 4, 5, 5)
	fmt.Println(r)

}

func add(x ...int) int {

	var total int

	for _, i := range x {
		total += i
	}

	// fmt.Println("The total sum of the nums are... ", total)
	// fmt.Print("The value of x is ", x)

	return total
}
