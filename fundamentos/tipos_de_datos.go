package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integer16 int16 = 16
	var integer32 int32 = 32

	s := "64"

	value, _ := strconv.Atoi(s)

	fmt.Println(value)
	fmt.Println(integer16 + int16(integer32))

	n := 8

	s = strconv.Itoa(n)
}
