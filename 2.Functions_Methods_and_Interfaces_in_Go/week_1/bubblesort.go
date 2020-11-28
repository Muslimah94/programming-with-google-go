8964package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter 10 numbers, or enter less and type 'X'")
	var input string
	var arr []int
	counter := 1
	for input != "X" {
		fmt.Println("An integer #", counter, ":")
		fmt.Scanln(&input)
		n, err := strconv.Atoi(input)
		if err != nil {
			if input != "X" {
				fmt.Println("Only valid integers, plz:", err)
				continue
			}
		} else {
			arr = append(arr, n)
		}
		if counter == 10 {
			break
		}
		counter++
	}
	BubbleSort(arr)
	fmt.Println(arr)
}

func Swap(numbers []int, i int) {
	tmp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = tmp
}

func BubbleSort(numbers []int) {
	Swapped := true
	for Swapped {
		Swapped = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				Swapped = true
			}
		}
	}
}
