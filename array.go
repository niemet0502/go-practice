package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var students []Student

	students = append(students, Student{"marius", 20})

	fmt.Println(students)
}
