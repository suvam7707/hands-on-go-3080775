// challenges/types-composite/begin/main.go
package main

// define an author type with a name
//
type author struct {
	first string
	last  string
}

// define a book type with a title and author
//
type book struct {
	title      string
	authorname author
}

// define a library type to track a list of books
//
type library struct {
	books []book
}

//var library = make([]book,0)
// define addBook to add a book to the library
//

// define a lookupByAuthor function to find books by an author's name
//

func main() {
	// create a new library
	//

	// add 2 books to the library by the same author
	//

	// add 1 book to the library by a different author
	//

	// dump the library with spew
	//

	// lookup books by known author in the library
	//

	// print out the first book's title and its author's name
	//

}
