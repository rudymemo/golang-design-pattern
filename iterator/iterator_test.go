package iterator

import (
	"testing"
)

func TestBookIterator(t *testing.T) {
	collection := NewBookCollection()
	collection.AddBook(Book{Title: "Go Programming", Author: "John Doe"})
	collection.AddBook(Book{Title: "Design Patterns", Author: "Gang of Four"})
	collection.AddBook(Book{Title: "Clean Code", Author: "Robert Martin"})

	iterator := collection.CreateIterator()
	count := 0

	for iterator.HasNext() {
		book := iterator.Next().(Book)
		if book.Title == "" {
			t.Error("Book title should not be empty")
		}
		count++
	}

	if count != 3 {
		t.Errorf("Expected 3 books, got %d", count)
	}
}

func TestNumberIterator(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	collection := NewNumberCollection(numbers)
	iterator := collection.CreateIterator()

	count := 0
	sum := 0

	for iterator.HasNext() {
		number := iterator.Next().(int)
		sum += number
		count++
	}

	if count != 5 {
		t.Errorf("Expected 5 numbers, got %d", count)
	}

	if sum != 15 {
		t.Errorf("Expected sum 15, got %d", sum)
	}
}