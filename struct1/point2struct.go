package main

import "learngo/struct1/book"

func main() {
	var Book1 book.Books /* Declare Book1 of type Book */
	var Book2 book.Books /* Declare Book2 of type Book */

	/* book 1 specification */
	Book1.Title = "Go Programming"
	Book1.Author = "Mahesh Kumar"
	Book1.Subject = "Go Programming Tutorial"
	Book1.Book_id = 6495407

	/* book 2 specification */
	Book2.Title = "Telecom Billing"
	Book2.Author = "Zara Ali"
	Book2.Subject = "Telecom Billing Tutorial"
	Book2.Book_id = 6495700

	/* print Book1 info */
	book.PrintBook(&Book1)

	/* print Book2 info */
	book.PrintBook(&Book2)

}
