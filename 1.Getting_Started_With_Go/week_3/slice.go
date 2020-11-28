package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Enter integers one by one and type 'X' letter when You are done")
	var input string
	slice := make([]int, 0, 3)
	for input != "X" {
		fmt.Scanln(&input)
		num, err := strconv.Atoi(input)
		if err != nil && input != "X" {
			fmt.Println("Only integers plz!", err.Error())
			continue
		}
		slice = append(slice, num)
		sort.Ints(slice)
		fmt.Println("Sorted slice of ints:", slice[1:])
	}

}
