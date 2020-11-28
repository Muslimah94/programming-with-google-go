package main

import (
	"fmt"
	"math/rand"
	"time"
)

var runes = []rune("一二三四五六七八九十1234567890")

func generateString(n int, str *string) {

	randRune := make([]rune, n)

	for i := range randRune {
		// without this, the final value will be same all the time.
		rand.Seed(time.Now().UnixNano())

		randRune[i] = runes[rand.Intn(len(runes))]
	}
	*str = string(randRune)

}
func main() {

	var str string
	go generateString(5, &str)
	fmt.Println(str)

	/* RACE CONDITION:
	Firstly, we have to understand that func main() in Go is the 1st and main goroutine.
	Second goroutine here is my func - generateString().

	both generateString and fmt.Println functions trying to access str string variable
	in order to write into and read from it. That's why output of execution is empty.

	If it's still not understandable, just erase "go" command on 27th line. And U will
	see that now there is no race condition and both funcs can write and read from this
	variable in a given order, that' why output won't be empty.*/

}
