package main

import "fmt"

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	streeingWheel int16
	topSpeedKmh   float64
}

func main() {
	Car := car{
		gasPedal:      222,
		brakePedal:    111,
		streeingWheel: 999,
		topSpeedKmh:   44.33,
	}

	fmt.Println(Car.brakePedal)
}
