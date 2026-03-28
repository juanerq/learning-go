package main

import "fmt"

func divide(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(num int) {
	if num == 0 {
		panic("No se puede dividir por 0")
	}
}

func main() {
	divide(10, 2)
	divide(18, 9)
	divide(10, 0)
	divide(9, 3)

}
