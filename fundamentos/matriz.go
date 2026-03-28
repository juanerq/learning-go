package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 4, 5}

	a[0] = 10

	for i := 0; i < len(a); i++ {
		a[i] = a[i] * 2
	}

	for index, values := range a {
		a[index] = values + 1
	}

	fmt.Print(a)
}
