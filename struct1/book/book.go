package book

import "fmt"

type Books struct {
	Title   string
	Author  string
	Subject string
	Book_id int
}

func PrintBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.Title)
	fmt.Printf("Book author : %s\n", book.Author)
	fmt.Printf("Book subject : %s\n", book.Subject)
	fmt.Printf("Book book_id : %d\n", book.Book_id)
}
