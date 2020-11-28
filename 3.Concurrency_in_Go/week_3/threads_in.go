package main

import (
	"fmt"
	"strconv"
	"sync"
)

func Swap(numbers []int, i int) {
	tmp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = tmp
}

func BubbleSort(wg *sync.WaitGroup, numbers []int) {
	fmt.Println(numbers)
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
	wg.Done()
}

func main() {
	fmt.Println("Enter integers to sort one by one.\nType 'X' if it's enough.")
	var input string
	var arr []int
	for input != "X" {
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
	}
	size := len(arr) / 4
	var wg sync.WaitGroup
	for c := 0; c < 4; c++ {
		wg.Add(1)
		if c != 3 {
			go BubbleSort(&wg, arr[c*size:(c+1)*size])
		} else {
			go BubbleSort(&wg, arr[c*size:])
		}
	}
	wg.Wait()

	var sorted []int
	c1 := arr[:1*size]
	c2 := arr[1*size : 2*size]
	c3 := arr[2*size : 3*size]
	c4 := arr[3*size:]
	for k := 0; k < len(arr); k++ {
		i := 0
		j := 0
		if len(c1) != 0 {
			i = c1[0]
			j = 1
		}
		if len(c2) != 0 {
			if j == 0 {
				i = c2[0]
				j = 2
			} else {
				if c2[0] < i {
					i = c2[0]
					j = 2
				}
			}
		}
		if len(c3) != 0 {
			if j == 0 {
				i = c3[0]
				j = 3
			} else {
				if c3[0] < i {
					i = c3[0]
					j = 3
				}
			}
		}
		if len(c4) != 0 {
			if j == 0 {
				i = c4[0]
				j = 4
			} else {
				if c4[0] < i {
					i = c4[0]
					j = 4
				}
			}
		}
		sorted = append(sorted, i)
		switch j {
		case 1:
			c1 = c1[1:]
		case 2:
			c2 = c2[1:]
		case 3:
			c3 = c3[1:]
		case 4:
			c4 = c4[1:]
		}
	}
	fmt.Println(sorted)
}
