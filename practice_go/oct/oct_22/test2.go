package main

import "fmt"

var name string = "Naruto"

func main() {

	const attack string = "Rasingan"
	say := "Hello"
	fmt.Println(say)
	fmt.Println("My name is", name)

	fmt.Println(attack)

	greet()

	spray()
}

func greet() {

	fmt.Println("Bye Bye", name)
}

func spray() {
	const (
		a = iota * 50
		b
		c
		d
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
