package main

import "time"

func main() {
	if t := time.Now(); t.Hour() < 12 {
		println("Buenos días")
	} else if t.Hour() < 18 {
		println("Buenas tardes")
	} else {
		println("Buenas noches")
	}
}
