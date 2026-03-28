package main

import (
	"errors"
	"fmt"
)

func dividir(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividor por 0")
	}
	return dividendo / divisor, nil
}

func main() {
	num, err := dividir(14, 0)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Println(num)

}
