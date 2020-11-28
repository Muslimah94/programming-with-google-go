package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter a random floating number, please:")
	var input string
	fmt.Scanln(&input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Only floating number is valid!", err.Error())
		os.Exit(1)
	}
	fmt.Println("Your input is converted from string to a float64:", num)
	fmt.Println("Now this floating number is converted to an integer:", int(num))
}
