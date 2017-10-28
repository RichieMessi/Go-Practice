package main

import (
	"fmt"
	"strconv"
)

// var foo float64 = 55.9

// var x = 10

func main() {

	var i int = 67
	fmt.Printf("%v, %T", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T", j, j)

	// Must use "strconv" import
	// Otherwise the compiler returns the ascii code of the conversion
	// Does not convert to string

}
