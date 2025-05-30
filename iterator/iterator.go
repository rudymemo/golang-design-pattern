package iterator

import "fmt"

// Iterator 迭代器接口
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Collection 集合接口
type Collection interface {
	CreateIterator() Iterator
}

// Book 书籍结构
type Book struct {
	Title  string
	Author string
}

func (b *Book) String() string {
	return fmt.Sprintf("%s by %s", b.Title, b.Author)
}

// BookCollection 书籍集合
type BookCollection struct {
	books []Book
}

func NewBookCollection() *BookCollection {
	return &BookCollection{
		books: make([]Book, 0),
	}
}

func (bc *BookCollection) AddBook(book Book) {
	bc.books = append(bc.books, book)
}

func (bc *BookCollection) CreateIterator() Iterator {
	return &BookIterator{
		collection: bc,
		index:      0,
	}
}

func (bc *BookCollection) GetBooks() []Book {
	return bc.books
}

// BookIterator 书籍迭代器
type BookIterator struct {
	collection *BookCollection
	index      int
}

func (bi *BookIterator) HasNext() bool {
	return bi.index < len(bi.collection.books)
}

func (bi *BookIterator) Next() interface{} {
	if bi.HasNext() {
		book := bi.collection.books[bi.index]
		bi.index++
		return book
	}
	return nil
}

// NumberCollection 数字集合（另一个例子）
type NumberCollection struct {
	numbers []int
}

func NewNumberCollection(numbers []int) *NumberCollection {
	return &NumberCollection{numbers: numbers}
}

func (nc *NumberCollection) CreateIterator() Iterator {
	return &NumberIterator{
		collection: nc,
		index:      0,
	}
}

// NumberIterator 数字迭代器
type NumberIterator struct {
	collection *NumberCollection
	index      int
}

func (ni *NumberIterator) HasNext() bool {
	return ni.index < len(ni.collection.numbers)
}

func (ni *NumberIterator) Next() interface{} {
	if ni.HasNext() {
		number := ni.collection.numbers[ni.index]
		ni.index++
		return number
	}
	return nil
}