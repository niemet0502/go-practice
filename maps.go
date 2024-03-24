package main

import "fmt"

type User struct {
	name string
	age  int
}

var m map[string]User

func main() {
	// initialize with an element
	m = map[string]User{
		"second": {"Lyam", 0},
	}

	// adding an element
	m["first"] = User{"marius", 24}

	fmt.Println(m["first"])
	fmt.Println(m)

	// deleting an element
	delete(m, "second")
	fmt.Println(m)

	_, ok := m["first"]

	fmt.Println(ok)
}
