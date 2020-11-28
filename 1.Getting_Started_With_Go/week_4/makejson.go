package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var input string
	m := make(map[string]string)
	fmt.Println("Enter Ur name:")
	fmt.Scanln(&input)
	m["name"] = input
	fmt.Println("Enter Ur address:")
	fmt.Scanln(&input)
	m["address"] = input
	fmt.Println("\nHere is ur data saved in a map:", m)
	jsonFile, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("\nAnd here is ur data marshalled in JSON object(it's a byte array):\n", jsonFile)
}
