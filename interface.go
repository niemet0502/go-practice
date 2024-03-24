package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vert struct {
	X, Y float64
}

func (v *Vert) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// implement implicitly an interface

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	fmt.Println("learning interface")

	v := Vert{1, 2}

	var a Abser

	a = &v

	fmt.Println(a.Abs())

}
