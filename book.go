package main

type (
	book struct {
		ID     uint
		Title  string
		Author string
	}

	bookModel interface {
		Create(b *book) error
		Read(id uint) (book, error)
		Update(b *book) error
		Delete(id uint) error
		List() []book
	}
)
