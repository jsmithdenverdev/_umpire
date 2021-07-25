package main

import (
	"context"
	"log"
)

type (
	getBookQuery struct {
		ID uint
	}

	getBookQueryHandler struct {
		logger *log.Logger
		model  bookModel
	}

	listBooksQuery struct {
	}

	listBooksQueryHandler struct {
		logger *log.Logger
		model  bookModel
	}
)

func (h *getBookQueryHandler) Handle(request interface{}, ctx context.Context) (interface{}, error) {
	query := request.(getBookQuery)

	h.logger.Printf("getting book with id %s", query.ID)

	return h.model.Read(query.ID)
}

func (h *listBooksQueryHandler) Handle(request interface{}, ctx context.Context) (interface{}, error) {
	h.logger.Printf("listing books")

	return h.model.List(), nil
}
