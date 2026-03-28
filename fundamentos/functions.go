package main

import "fmt"

func main() {
	saludo := hello("Juan")
	fmt.Println(saludo)

	result := add(8, 5)
	fmt.Println(result)
}

func hello(name string) string {
	return fmt.Sprintf("Hola %s\n", name)
}

func add(a, b int) int {
	return a + b
}

func cal(a, b int) (sum, diff int) {
	sum = a + b
	diff = a - b
	return
}
