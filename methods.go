package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) display() string {
	return fmt.Sprintf("I'm %s i'm %d years old", p.name, p.age)
}

func (p *Person) update(name string, age int) {
	p.name = name
	p.age = age
}

func main() {
	p := Person{"marius", 20}

	fmt.Println(p.display())

	p.update("niemet marius", 50)

	fmt.Println(p.display())

}
