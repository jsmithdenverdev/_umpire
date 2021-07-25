package main

import (
	"context"
	"fmt"
	"reflect"
)

type mediator struct {
	requestHandlers       map[string]requestHandler
	requestHandlersReturn map[string]requestHandlerReturn
}

func newMediator() *mediator {
	requestHandlers := make(map[string]requestHandler, 0)
	requestHandlersReturn := make(map[string]requestHandlerReturn, 0)

	return &mediator{
		requestHandlers:       requestHandlers,
		requestHandlersReturn: requestHandlersReturn,
	}
}

// registerRequestHandler registers a given request with a requestHandler. A request may have a single requestHandler.
// An error is returned if the given request has already been registered.
func (m *mediator) registerRequestHandler(request interface{}, handler requestHandler) error {
	requestName := reflect.TypeOf(request).Name()

	if handler, ok := m.requestHandlers[requestName]; ok {
		handlerName := reflect.TypeOf(handler).Name()

		return fmt.Errorf("request %s already has registered handler %s", requestName, handlerName)
	}

	m.requestHandlers[requestName] = handler

	return nil
}

// registerRequestHandlerReturn registers a given request with a requestHandlerReturn. A request may have a single
// requestHandlerReturn. An error is returned if the given request has already been registered.
func (m *mediator) registerRequestHandlerReturn(request interface{}, handler requestHandlerReturn) error {
	requestName := reflect.TypeOf(request).Name()

	if handler, ok := m.requestHandlersReturn[requestName]; ok {
		handlerName := reflect.TypeOf(handler).Name()

		return fmt.Errorf("request %s already has registered handler %s", requestName, handlerName)
	}

	m.requestHandlersReturn[requestName] = handler

	return nil
}

// dispatch attempts to find and execute a requestHandler for the given request returning an error if the request has no
// registered requestHandler. dispatch will bubble the error from the requestHandler.
func (m *mediator) dispatch(request interface{}, ctx context.Context) error {
	requestName := reflect.TypeOf(request).Name()

	if handler, ok := m.requestHandlers[requestName]; ok {
		return handler.Handle(request, ctx)

	}

	return fmt.Errorf("request %s has no handlers", requestName)
}

// dispatchReturn attempts to find and execute a requestHandlerReturn for the given request returning an error if the
// request has no registered requestHandlerReturn. dispatch will bubble the result and error from the
// requestHandlerReturn.
func (m *mediator) dispatchReturn(request interface{}, ctx context.Context) (interface{}, error) {
	requestName := reflect.TypeOf(request).Name()

	if handler, ok := m.requestHandlersReturn[requestName]; ok {
		return handler.Handle(request, ctx)

	}

	return nil, fmt.Errorf("request %s has no handlers", requestName)
}
