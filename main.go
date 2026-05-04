package main

import (
	"golang/poo/book"
)

func main() {
	var myBook = book.NewBook("The Go Programming Language", "Alan A. A. Donovan", 380)

	myBook.PrintInfo()
}
