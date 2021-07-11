package main

import (
	"fmt"
	"iterator/book"
)

func main() {
	bs := book.NewBookShelf(4)
	bs.AppendBook(book.NewBook("Around the World in 80 Days"))
	bs.AppendBook(book.NewBook("Bible"))
	bs.AppendBook(book.NewBook("Cinderella"))
	bs.AppendBook(book.NewBook("Daddy-Long-Legs"))

	it := bs.CreateIterator()
	for it.HasNext() {
		bk := it.Next()
		fmt.Println(bk.GetName())
	}

}
