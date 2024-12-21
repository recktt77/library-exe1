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
	_, ok := l.Books[id]
	if !ok {
		fmt.Printf("Book doesn't exist")
	}
	l.Books[id].isBorrowed = true
	fmt.Printf("Book '%s' succesfully borrowed", l.Books[id].title)
}

func (l *Library) ReturnBook(id string) {
	_, ok := l.Books[id]
	if !ok {
		fmt.Printf("This book is not from this library")
	}
	l.Books[id].isBorrowed = false
	fmt.Printf("Book '%s' succesfully returned", l.Books[id].title)
}

func (l *Library) DisplayBooks() {
	for _, value := range l.Books {
		fmt.Printf("'%s' '%s' '%s' '%v' \n", value.id, value.title, value.author, value.isBorrowed)
	}
}
