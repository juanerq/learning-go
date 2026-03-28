package main

import (
	"fmt"
)

const PI float32 = 3.1416

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func variablas() {
	var fistName string = "Juan"

	var (
		lastName string = "Reyes"
		age      int    = 24
	)

	fmt.Println(Viernes)
	fmt.Println("Hola " + fistName + " " + lastName + ", tienes " + fmt.Sprint(age) + " años.")
}
