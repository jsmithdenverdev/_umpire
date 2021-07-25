package main

import (
	"fmt"
)

type mockBookModel struct {
	last  book
	books map[uint]book
}

func newMockBookModel() *mockBookModel {
	books := make(map[uint]book, 0)

	return &mockBookModel{
		book{},
		books,
	}
}

func (m *mockBookModel) Create(b *book) error {
	b.ID = m.last.ID + 1

	if _, ok := m.books[b.ID]; ok {
		return fmt.Errorf("duplicate identifier found %d", b.ID)
	}

	m.books[b.ID] = *b
	m.last = *b

	return nil
}

func (m *mockBookModel) Read(id uint) (book, error) {
	b, ok := m.books[id]

	if ok {
		return b, nil
	}

	return book{}, fmt.Errorf("no book found with identifier %d", b.ID)
}

func (m *mockBookModel) Update(b *book) error {
	if book, ok := m.books[b.ID]; ok {
		book.Title = b.Title
		book.Author = b.Author

		m.books[book.ID] = book

		return nil
	}

	return fmt.Errorf("tried updating non-existant book with identifier %d", b.ID)
}

func (m *mockBookModel) Delete(id uint) error {
	if _, ok := m.books[id]; ok {
		delete(m.books, id)
	}

	return fmt.Errorf("tried deleting non-existant book with identifier %d", id)
}

func (m *mockBookModel) List() []book {
	books := make([]book, 0)

	for _, b := range m.books {
		books = append(books, b)
	}

	return books
}
