package library

import (
	"fmt"
	"strconv"
)

type Book struct {
	id         string
	title      string
	author     string
	isBorrowed bool
}

type Library struct {
	Books  map[string]*Book
	nextID int
}

func (l *Library) Init() {
	if l.Books == nil {
		l.Books = make(map[string]*Book)
	}
}

func (l *Library) AddBooks(title, author string) {
	id := strconv.Itoa(l.nextID)
	l.Books[id] = &Book{
		id:         id,
		title:      title,
		author:     author,
		isBorrowed: false,
	}
	l.nextID++
	fmt.Printf("book '%s' added", title)
}

func (l *Library) BorrowBook(id string) {
	book, ok := l.Books[id]
	if !ok {
		fmt.Println("Book doesn't exist.")
		return
	}
	if book.isBorrowed {
		fmt.Printf("Book '%s' is already borrowed.\n", book.title)
		return
	}
	book.isBorrowed = true
	fmt.Printf("Book '%s' successfully borrowed.\n", book.title)
}

func (l *Library) ReturnBook(id string) {
	book, ok := l.Books[id]
	if !ok {
		fmt.Println("This book is not from this library.")
		return
	}
	if !book.isBorrowed {
		fmt.Printf("Book '%s' is already returned.\n", book.title)
		return
	}
	book.isBorrowed = false
	fmt.Printf("Book '%s' successfully returned.\n", book.title)
}

func (l *Library) DisplayBooks() {
	for _, value := range l.Books {
		if !value.isBorrowed{
			fmt.Printf("'%s' '%s' '%s' '%v' \n", value.id, value.title, value.author, value.isBorrowed)
		}
	}
}
