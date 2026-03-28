package main

import "fmt"

type Person struct {
	name  string
	age   int
	email string
}

func main() {
	person1 := Person{
		name:  "Alice",
		age:   30,
		email: "alice@example.com",
	}

	fmt.Print(person1)
}
