package main

import (
	"context"
	"log"
)

type (
	createBookCommand struct {
		Title  string
		Author string
	}

	createBookCommandHandler struct {
		logger *log.Logger
		model  bookModel
	}
)

func (h *createBookCommandHandler) Handle(request interface{}, ctx context.Context) (interface{}, error) {
	command := request.(createBookCommand)

	b := book{
		Title:  command.Title,
		Author: command.Title,
	}

	h.logger.Printf("creating book %+v", b)

	if err := h.model.Create(&b); err != nil {
		return book{}, err
	}

	h.logger.Printf("created book %+v", b)

	return b, nil
}
