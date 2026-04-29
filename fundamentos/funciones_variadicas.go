package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

func printData(data ...interface{}) {
	for _, item := range data {
		fmt.Printf("Valor: %v, Tipo: %T\n", item, item)
	}
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	println(result)

	printData("Hello", 123, true, 45.67)
}
