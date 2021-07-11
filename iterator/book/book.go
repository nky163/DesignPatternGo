package book

type Iterator interface {
	HasNext() bool
	Next() Book // ...
}

type Aggregate interface {
	CreateIterator() Iterator
}

type BookShelf struct {
	books []*Book
	last  int64
}

func NewBookShelf(maxSize uint64) *BookShelf {
	b := BookShelf{}
	b.books = make([]*Book, maxSize, maxSize)
	b.last = 0
	return &b
}

func (bs *BookShelf) CreateIterator() Iterator {
	return &BookShelfIterator{bs, 0}
}

func (bs *BookShelf) GetBookAt(index int64) *Book {
	if index > int64(len(bs.books)) {
		return nil
	}
	return bs.books[index]
}

func (bs *BookShelf) GetLength() int64 {
	return int64(len(bs.books))
}

func (bs *BookShelf) AppendBook(book *Book) {
	if bs.last >= int64(cap(bs.books)) {
		bs.books = append(bs.books, book)
	}
	bs.books[bs.last] = book
	bs.last++
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int64
}

func (bsi *BookShelfIterator) HasNext() bool {
	if bsi.index == bsi.bookShelf.last {
		bsi.index = 0
		return false
	}
	return true
}

func (bsi *BookShelfIterator) Next() Book {
	b := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index++
	return *b
}

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name}
}

func (b *Book) GetName() string {
	return b.name
}
