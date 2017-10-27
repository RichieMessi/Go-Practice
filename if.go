package main

import "fmt"

func main() {

	age := 17

	if age > 1 {
		fmt.Println("you can drink in the club...")
	} else if age == 17 {
		fmt.Println("Naa.. you are 17... Go finish your homework  and come back in one year")
	} else {
		fmt.Println("you are not eligible ")
	}

}
