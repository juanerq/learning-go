package main

import "fmt"

type Animal struct {
	name string
	age  int
	size float32
}

func (a *Animal) greet() {
	fmt.Printf("Hola soy %s", a.name)
}

func main() {
	var x int = 10
	var p *int = &x
	editValue(p, 7)

	cat := Animal{"Newton", 10, 2}
	cat.greet()
}

func editValue(value *int, newValue int) {
	*value = newValue
}
