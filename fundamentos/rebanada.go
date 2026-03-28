package main

import "fmt"

func main() {
	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// daysCopy := days[0:5]

	// daysCopy = append(daysCopy, "Lunes", "Martes")

	// daysCopy2 := append(days[:2], days[4:]...)

	// names := make([]string, 5, 10)

	// fmt.Println(cap((daysCopy)))

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)

	copy(slice2, slice1)

	fmt.Println(slice2)
}
