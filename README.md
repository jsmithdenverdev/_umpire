# _umpire

A naive and incomplete mediator example in Go.

## Motivation

I wanted to see what the feasibility of a mediator implementation with an API similar
to [mediatr](https://github.com/jbogard/MediatR) but written in pure Golang was. Written in a couple of hours with a couple of 🥃 on a Saturday night.

## Example usage

```go
// create a new request and handler, this request returns a result
type (
	createBookCommand struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	createBookCommandHandler struct {
		model bookModel
	}
)

// implement requestReturnHandler for createBookCommandHandler
func (h *createBookCommandHandler) Handle(request interface{}, ctx context.Context) (interface{}, error) {
	command := request.(createBookCommand)

	b := book{
		command.Title,
		command.Author,
	}

	return h.model.Create(&b)
}

// -----------------------

func main() {
	// create a new mediator
	umpire := newMediator()
	model := newMockModel()

	// register the request with its appropriate handler
	umpire.registerReturnHandler(createBookCommand{}, &createBookCommandHandler{
		model,
	})

	// tell the mediator to dispatch the given command and return the result (and error)
	book, err := umpire.dispatchReturn(createBookCommand{
		Title:  "Jake's Book",
		Author: "Jake Smith",
	})

	if err != nil {
		panic(err)
	}

	// print the created book
	fmt.Printf("Created book %+v", book)
}
```

## Issues
This pattern suffers a major flaw with Go in its current state (pre-generics).

The `requestHandler` and `requestReturnHandler` interfaces each take their respective requests as the first argument to their `Handle` method.  Without generics, I must write the type signature of `requestHandler` as `Handle(request interface{}, ctx context.Context) error` and `requestReturnHandler` as `Handle(request interface{}, ctx context.Context) (interface{}, error)`. This forces every handler to perform a type assertion to convert the `interface{}` into the type it expects.  This also means that `requestReturnHandler`'s return an `interface{}` which drastically decreases type safety and requires callers to perform another type assertion to convert the `interface{}` into the correct type.