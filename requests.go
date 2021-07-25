package main

import "context"

type (
	requestHandler interface {
		Handle(request interface{}, ctx context.Context) error
	}

	requestReturnHandler interface {
		Handle(request interface{}, ctx context.Context) (interface{}, error)
	}
)
