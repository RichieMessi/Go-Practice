package main

import "fmt"

const usixteenbitmax float64 = 5555
const kmhMultiple float64 = 1.234

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	streeingWheel int16
	topSpeedKmh   float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func main() {
	Car := car{
		gasPedal:      222,
		brakePedal:    111,
		streeingWheel: 999,
		topSpeedKmh:   44.33,
	}

	fmt.Println(Car.brakePedal)
	fmt.Println(Car.kmh())
}
