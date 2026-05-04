package book

import "fmt"

type book struct {
	title  string
	author string
	pages  int
}

func NewBook(title, author string, pages int) *book {
	return &book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *book) PrintInfo() {
	fmt.Printf("Tile: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}
