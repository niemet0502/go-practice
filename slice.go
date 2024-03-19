package main

import "fmt"

func main() {
	// slice is initliaze with NIL
	var m []int
	fmt.Println(m, len(m), cap(m))

	// Array with a fixed size of 3
	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slice without a fixed size
	slice := []int{4, 5, 6}
	fmt.Println("Slice:", slice)

	// the first is excluded
	slice2 := slice[0:2]
	fmt.Println("Slice 2 :", slice2)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	// len(s) => slice length
	// cap(s) => slice capacity

	fmt.Println("%d length and %d capacity", len(s), cap(s))

	// create a slice by using the built-in function make
	t := make([]int, 5, 10)
	fmt.Println(t)

	// add a new value

	t = append(t, 44)
	fmt.Println(t)
}
