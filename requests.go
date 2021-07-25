package main

import "context"

type (
	requestHandler interface {
		Handle(request interface{}, ctx context.Context) error
	}

	requestHandlerReturn interface {
		Handle(request interface{}, ctx context.Context) (interface{}, error)
	}
)
