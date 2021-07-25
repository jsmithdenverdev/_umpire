package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func main() {
	model := newMockBookModel()
	umpire := newMediator()
	logger := log.New(os.Stdout, "INFO ", log.Ldate)
	ctx := context.Background()

	var (
		commands       = map[interface{}]requestHandler{}
		commandsReturn = map[interface{}]requestHandlerReturn{
			createBookCommand{}: &createBookCommandHandler{
				logger,
				model,
			},
		}
		queries       = map[interface{}]requestHandler{}
		queriesReturn = map[interface{}]requestHandlerReturn{
			getBookQuery{}: &getBookQueryHandler{
				logger,
				model,
			},
			listBooksQuery{}: &listBooksQueryHandler{
				logger,
				model,
			},
		}
	)

	for command, handler := range commands {
		if err := umpire.registerRequestHandler(command, handler); err != nil {
			logger.Fatal(err)
		}
	}

	for query, handler := range queries {
		if err := umpire.registerRequestHandler(query, handler); err != nil {
			logger.Fatal(err)
		}
	}

	for command, handler := range commandsReturn {
		if err := umpire.registerRequestHandlerReturn(command, handler); err != nil {
			logger.Fatal(err)
		}
	}

	for query, handler := range queriesReturn {
		if err := umpire.registerRequestHandlerReturn(query, handler); err != nil {
			logger.Fatal(err)
		}
	}

	for i := 0; i < 10; i++ {
		command := createBookCommand{
			Title:  fmt.Sprintf("Book %d", i),
			Author: fmt.Sprintf("Author %d", i),
		}

		if _, err := umpire.dispatchReturn(command, ctx); err != nil {
			logger.Fatal(err)
		}
	}

	books, err := umpire.dispatchReturn(listBooksQuery{}, ctx)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Books %+v", books)
}
