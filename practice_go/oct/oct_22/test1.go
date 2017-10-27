package main

import "fmt"

func main() {

	fmt.Println("Hello")

	mylist := []int{1, 2, 3, 4, 5}

	fmt.Println(mylist)

	fmt.Println("Sum :", add(mylist))

}

func add(nums []int) int {

	sum := 0

	for _, val := range nums {

		sum += val
	}

	return sum
}

/* space */
