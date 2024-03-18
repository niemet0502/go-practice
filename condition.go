package main

import (
	"fmt"
	"strconv"
)

func calculate(x, y int) string {
	// the v variable is only available inside the condition
	if v := x * 2; v > y {
		return "yes"
	} else {
		return strconv.Itoa(v) + " failed the test"
	}
}

func main() {
	age := 10

	number := 22

	if number > 18 {
		fmt.Println("older")
	}

	if age < 18 {
		fmt.Println("younger")
	}

	fmt.Println(calculate(1, 5))
}
