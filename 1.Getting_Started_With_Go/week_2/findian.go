package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Type some random string, please:")
	var input string
	fmt.Scanln(&input)
	str := fmt.Sprintf("%s", input)
	if strings.ToLower(string(str[0])) == "i" && strings.ToLower(string(str[len(str)-1])) == "n" && strings.Contains(str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
