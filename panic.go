package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// read the whole file at once
	b, err := ioutil.ReadFile("names.txt")
	if err == nil {
		panic("fuck theres was an error. Save women, children and gather food ")
	}

	for _, e := range b {
		fmt.Print(string(e))
	}
}
