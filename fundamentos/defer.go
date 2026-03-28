package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hola.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err = file.Write([]byte("hola Juan")); err != nil {
		fmt.Println(err)
		return
	}

}
