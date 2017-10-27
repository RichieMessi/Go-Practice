package main

import "fmt"

func main() {

	var arr [5]float64

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println(arr)
	fmt.Println("Array at index 3 is ", arr[2])

	arrCombine := [5]float64{1, 2, 3, 4, 5}

	fmt.Println(arrCombine)

	for _, value := range arrCombine {
		fmt.Println(value)
	}

}
