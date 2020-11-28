package main

import "fmt"

func main() {
	fmt.Println("Please input 3 values one by one\n-------------------------------")
	var acceleration float64
	fmt.Println("Acceleration:")
	fmt.Scanln(&acceleration)
	var displacement float64
	fmt.Println("Initial displacement:")
	fmt.Scanln(&displacement)
	var velocity float64
	fmt.Println("Initial velocity:")
	fmt.Scanln(&velocity)
	fn := GenDisplaceFn(acceleration, displacement, velocity)
	var time float64
	fmt.Println("Now enter time:")
	fmt.Scanln(&time)

	fmt.Println("RESULT:", fn(time))
}

func GenDisplaceFn(acceleration, displacement, velocity float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return acceleration*time*time/2 + velocity*time + displacement
	}
	return fn
}
