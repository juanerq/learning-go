package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Ejecutando en windows")
	case "linux":
		fmt.Println("Ejecutando en linux")
	default:
		fmt.Println("Ejecutando en otro sistema operativo")
	}
}
