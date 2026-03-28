package main

import "fmt"

func main() {
	var name string
	var lastName string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name, &lastName)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	greeting := fmt.Sprintf("Hola, me llamo %s %s y tengo %d años.", name, lastName, age)
	fmt.Println(greeting)

	fmt.Printf("%T", name)
}
